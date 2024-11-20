package loketrepository

import (
	"be-mklinik/entities"

	"gorm.io/gorm"
)

type LoketRepositoryInterface interface {
	GetDataLoket() ([]entities.MLoket, error)
}

type connection struct {
	db *gorm.DB
}

func NewLoketRepository(dbParam *gorm.DB) *connection {
	return &connection{
		db: dbParam,
	}
}

func (c *connection) GetDataLoket() ([]entities.MLoket, error) {
	var lokets []entities.MLoket
	result := c.db.Where("aktif = ?", 1).Find(&lokets)
	return lokets, result.Error
}
