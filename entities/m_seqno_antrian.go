package entities

import "time"

type MSeqNoAntrian struct {
	Id                  string    `json:"id" gorm:"column:id"`
	Tanggal             time.Time `json:"tanggal" gorm:"column:tanggal"`
	KodeAntrianKategori string    `json:"kode_antrian_kategori" gorm:"column:kode_antrian_kategori"`
	Seqno               int       `json:"seqno" gorm:"column:seqno"`
	Status              int       `json:"status" gorm:"column:status"`
	Panggil             int       `json:"panggil" gorm:"column:panggil"`
	Aktif               int       `json:"aktif" gorm:"column:aktif"`
	KodeLoket           string    `json:"kode_loket" gorm:"column:kode_loket"`
	CreatedAt           time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (MSeqNoAntrian) TableName() string {
	return "m_seqno_antrian"
}
