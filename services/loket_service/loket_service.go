package loketservice

import (
	"be-mklinik/entities"
	loketrepository "be-mklinik/repositories/loket_repository.go"
	"be-mklinik/responses"
	"fmt"
)

type LoketServiceInterface interface {
	GetDataLoket() ([]responses.LoketResponse, error)
}

type loketRepository struct {
	repository loketrepository.LoketRepositoryInterface
}

func NewLoketService(repository loketrepository.LoketRepositoryInterface) *loketRepository {
	return &loketRepository{repository: repository}
}
func (s *loketRepository) GetDataLoket() ([]responses.LoketResponse, error) {
	// Channel untuk komunikasi antar goroutine
	errorChan := make(chan error, 1)
	dataLoketChan := make(chan []entities.MLoket, 1)
	dataResponseChan := make(chan []responses.LoketResponse, 1)

	// Goroutine untuk mengambil data loket
	go func() {

		dataLoket, err := s.repository.GetDataLoket()
		if err != nil {
			errorChan <- fmt.Errorf("error fetching data loket: %w", err)
			return
		}

		dataLoketChan <- dataLoket
	}()

	// Goroutine untuk memproses data loket menjadi response
	go func() {

		dataLoket, ok := <-dataLoketChan
		if !ok {
			// Channel sudah ditutup tanpa data
			errorChan <- fmt.Errorf("failed to fetch data loket")
			return
		}

		var dataLoketResponse []responses.LoketResponse
		for _, item := range dataLoket {
			dataLoketResponse = append(dataLoketResponse, responses.LoketResponse{
				Id:        item.Id,
				Kode:      item.Kode,
				Deskripsi: item.Deskripsi,
				Aktif:     item.Aktif,
			})
		}

		dataResponseChan <- dataLoketResponse
	}()

	// Select untuk menangani hasil atau error
	select {
	case err := <-errorChan:
		defer close(errorChan)
		return nil, err
	case result := <-dataResponseChan:
		defer close(dataResponseChan)
		defer close(dataLoketChan)
		return result, nil
	}
}

//
