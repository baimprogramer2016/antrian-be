package entities

import "time"

type MLoket struct {
	Id        string    `json:"id" gorm:"column:id"`
	Kode      string    `json:"kode" gorm:"column:kode"`
	Deskripsi string    `json:"Deskripsi" gorm:"column:deskripsi"`
	Aktif     int       `json:"aktif" gorm:"column:aktif"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (MLoket) TableName() string {
	return "m_loket"
}
