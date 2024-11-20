package monitorantrianservice

import (
	"be-mklinik/entities"
	monitorantrianrepository "be-mklinik/repositories/monitor_antrian_repository"
	"be-mklinik/requests"
	"be-mklinik/responses"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type MonitorAntrianService interface {
	GetAllMSeqnoAntrianByDay() (responses.MonitorAntrianResponse, error)
	UpdatePanggilanAntrian(request requests.ParamUpdateAntrianRequest) (entities.MSeqNoAntrian, error)
}

type monitorAntrianServiceRespository struct {
	repository monitorantrianrepository.MonitorAntrianRepository
}

func NewMonitorAntrianService(repository monitorantrianrepository.MonitorAntrianRepository) *monitorAntrianServiceRespository {
	return &monitorAntrianServiceRespository{repository: repository}
}

func (s *monitorAntrianServiceRespository) GetAllMSeqnoAntrianByDay() (responses.MonitorAntrianResponse, error) {

	var wg sync.WaitGroup
	dataKunjunganChan := make(chan []entities.MSeqNoAntrian)
	kategoriAntrianChan := make(chan string)
	seqnoChan := make(chan int)
	nomorPanggilChan := make(chan string)
	loketChan := make(chan string)
	textPanggilanChan := make(chan string)
	errorChan := make(chan error)
	resultJson := make(chan responses.MonitorAntrianResponse)

	wg.Add(2)
	go func() {

		defer wg.Done()

		dataKunjungan, err := s.repository.GetAllMSeqnoAntrianByDay()

		if err != nil {
			errorChan <- fmt.Errorf("error fetching data jumlah kunjungan: %w", err)
			return
		}

		//nomor panggilan
		kode_loket := ""
		nomor_panggil := ""
		kategori_antrian := ""
		seqno_antrian := 0
		for _, item := range dataKunjungan {
			if item.Panggil == 1 {
				nomor_panggil = item.KodeAntrianKategori + strconv.Itoa(item.Seqno)
				kode_loket = item.KodeLoket
				kategori_antrian = item.KodeAntrianKategori
				seqno_antrian = item.Seqno
			}
		}

		dataKunjunganChan <- dataKunjungan
		kategoriAntrianChan <- kategori_antrian
		seqnoChan <- seqno_antrian
		nomorPanggilChan <- nomor_panggil
		loketChan <- kode_loket
		textPanggilanChan <- strings.Join(strings.Split(nomor_panggil, ""), " ")

	}()

	go func() {
		defer wg.Done()
		// //LOKET AKTIF
		dataLoket, err := s.repository.GetAllMLoket()
		if err != nil {
			errorChan <- fmt.Errorf("error fetching Loket: %w", err)
			return
		}
		dataKunjungan := <-dataKunjunganChan
		kategoriAntrian := <-kategoriAntrianChan
		seqno := <-seqnoChan
		nomorPanggil := <-nomorPanggilChan
		kodeLoket := <-loketChan
		textPanggilan := <-textPanggilanChan

		//detail antrian
		var antrian []responses.MonitorAntrianSaatIniResponse

		for _, item_loket := range dataLoket {

			//default

			var kategori_antrian_item string = ""
			var seqno_item int = 0
			var nomor_antrian_item string = ""
			var aktif_item int = 0
			var panggil_item int = 0

			//ubah jika sedang dilayani
			for _, item_kunjugan := range dataKunjungan {
				if item_kunjugan.KodeLoket == item_loket.Kode && item_kunjugan.Aktif == 1 {
					kategori_antrian_item = item_kunjugan.KodeAntrianKategori
					seqno_item = item_kunjugan.Seqno
					nomor_antrian_item = fmt.Sprintf("%s%s", item_kunjugan.KodeAntrianKategori, strconv.Itoa(item_kunjugan.Seqno))
					aktif_item = item_kunjugan.Aktif
					panggil_item = item_kunjugan.Panggil

				}
			}
			antrian = append(antrian, responses.MonitorAntrianSaatIniResponse{
				Loket:           item_loket.Kode,
				NomorAntrian:    nomor_antrian_item,
				Aktif:           aktif_item,
				Panggil:         panggil_item,
				KategoriAntrian: kategori_antrian_item,
				Seqno:           seqno_item,
			})
		}
		//FINAL
		var result = responses.MonitorAntrianResponse{
			JumlahKunjungan: len(dataKunjungan),
			KategoriAntrian: kategoriAntrian,
			Seqno:           seqno,
			NomorPanggil:    nomorPanggil,
			Loket:           kodeLoket,
			TextPanggilan:   fmt.Sprintf("NOMOR ANTRIAN %s MENUJU LOKET %s", textPanggilan, kodeLoket),
			DataAntrian:     antrian,
		}

		resultJson <- result

	}()

	select {
	case err := <-errorChan:
		return responses.MonitorAntrianResponse{}, err
	case result := <-resultJson:
		wg.Wait()
		defer close(dataKunjunganChan)
		defer close(kategoriAntrianChan)
		defer close(seqnoChan)
		defer close(nomorPanggilChan)
		defer close(loketChan)
		defer close(textPanggilanChan)
		defer close(errorChan)
		return result, nil
	}
}

func (s *monitorAntrianServiceRespository) UpdatePanggilanAntrian(request requests.ParamUpdateAntrianRequest) (entities.MSeqNoAntrian, error) {

	errorChan := make(chan error)
	dataResponse := make(chan entities.MSeqNoAntrian)

	go func() {
		_, err := s.repository.UpdatePanggilanAntrian(request)
		if err != nil {
			errorChan <- err
			return
		}
	}()

	go func() {
		result, err := s.repository.UpdateAktifLoket(request)
		if err != nil {
			errorChan <- err
			return
		}

		dataResponse <- result

	}()

	select {
	case err := <-errorChan:
		return entities.MSeqNoAntrian{}, err
	case result := <-dataResponse:
		defer close(dataResponse)
		return result, nil
	}
}
