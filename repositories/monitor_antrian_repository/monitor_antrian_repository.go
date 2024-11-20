package monitorantrianrepository

import (
	"be-mklinik/entities"
	"be-mklinik/requests"
	"time"

	"gorm.io/gorm"
)

type MonitorAntrianRepository interface {
	GetAllMSeqnoAntrianByDay() ([]entities.MSeqNoAntrian, error)
	GetAllMLoket() ([]entities.MLoket, error)
	UpdatePanggilanAntrian(requestUpdate requests.ParamUpdateAntrianRequest) (entities.MSeqNoAntrian, error)
	UpdateAktifLoket(requestUpdate requests.ParamUpdateAntrianRequest) (entities.MSeqNoAntrian, error)
}

type connection struct {
	db *gorm.DB
}

func NewMonitorAntrianRepository(dbParam *gorm.DB) *connection {
	return &connection{db: dbParam}
}

func (conn *connection) GetAllMSeqnoAntrianByDay() ([]entities.MSeqNoAntrian, error) {
	var mseqnoantrian []entities.MSeqNoAntrian
	result := conn.db.Where("tanggal = ?", time.Now().Format("2006-01-02")).Find(&mseqnoantrian)

	return mseqnoantrian, result.Error
}

func (conn *connection) GetAllMLoket() ([]entities.MLoket, error) {
	var mloket []entities.MLoket
	result := conn.db.Where("aktif = ?", 1).Find(&mloket)
	return mloket, result.Error
}

//PANGGILAN
// --kolom panggil dijadikan 0 semua lalu update jadi 1 sesuai kode_antrian+seqno
// --kolom kode loket di update sesuai loket yang manggil berdasarkan kode_antrian+seqno
// --aktif dijadikan 0 untuk berdasarakan loket yang di pilih lalu dijadikan 1 berdasarakan kode_antrian+seqno

func (conn *connection) GetSeqnoAntrianById(id string) (entities.MSeqNoAntrian, error) {
	var mseqnoantrian entities.MSeqNoAntrian
	result := conn.db.First(&mseqnoantrian, id)
	return mseqnoantrian, result.Error
}

// db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
func (conn *connection) UpdatePanggilanAntrian(requestUpdate requests.ParamUpdateAntrianRequest) (entities.MSeqNoAntrian, error) {

	var updateField = entities.MSeqNoAntrian{
		Panggil: 0,
	}
	//update batch pada hari ini jadikan 0
	result := conn.db.Model(&entities.MSeqNoAntrian{}).Where("tanggal = ?", time.Now().Format("2006-01-02")).Select("panggil").Updates(updateField)

	if result.Error != nil {
		return entities.MSeqNoAntrian{}, result.Error
	}
	//update batch pada hari ini yang dipanggil saja jadi 1
	var dataUpdate entities.MSeqNoAntrian
	resultUpdate := conn.db.Where("id = ?", requestUpdate.Id).First(&dataUpdate)

	if resultUpdate.Error != nil {
		return entities.MSeqNoAntrian{}, resultUpdate.Error
	}
	// fmt.Println("dataUpdate", dataUpdate)
	var updateFieldPanggil = entities.MSeqNoAntrian{
		Panggil:   1,
		KodeLoket: requestUpdate.KodeLoket,
	}
	err := conn.db.Model(&entities.MSeqNoAntrian{}).Where("id = ?", requestUpdate.Id).Updates(updateFieldPanggil)
	if err.Error != nil {
		return entities.MSeqNoAntrian{}, err.Error
	}

	return dataUpdate, nil
}

func (conn *connection) UpdateAktifLoket(requestUpdate requests.ParamUpdateAntrianRequest) (entities.MSeqNoAntrian, error) {

	var updateField = entities.MSeqNoAntrian{
		Aktif: 0,
	}
	result := conn.db.Model(&entities.MSeqNoAntrian{}).Where("tanggal = ? AND kode_loket = ?", time.Now().Format("2006-01-02"), requestUpdate.KodeLoket).Select("aktif").Updates(updateField)

	if result.Error != nil {
		return entities.MSeqNoAntrian{}, result.Error
	}
	//update batch pada hari ini yang dipanggil saja jadi 1
	var dataUpdate entities.MSeqNoAntrian
	resultUpdate := conn.db.Where("id = ?", requestUpdate.Id).First(&dataUpdate)

	if resultUpdate.Error != nil {
		return entities.MSeqNoAntrian{}, resultUpdate.Error
	}
	// fmt.Println("dataUpdate", dataUpdate)
	var updateFieldPanggil = entities.MSeqNoAntrian{
		Aktif:     1,
		KodeLoket: requestUpdate.KodeLoket,
	}
	err := conn.db.Model(&entities.MSeqNoAntrian{}).Where("id = ?", requestUpdate.Id).Updates(updateFieldPanggil)
	if err.Error != nil {
		return entities.MSeqNoAntrian{}, err.Error
	}

	return dataUpdate, nil

}
