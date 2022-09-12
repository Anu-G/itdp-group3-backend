package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type BusinessHourRepositoryInterface interface {
	Create(hours *entity.BusinessHour) error
	Delete(id string) error
}

type businessHourRepository struct {
	db *gorm.DB
}

func (b *businessHourRepository) Create(hours *entity.BusinessHour) error {
	return b.db.Create(&hours).Error
}

func (b *businessHourRepository) Delete(id string) error {
	return b.db.Where("id = ?", id).Delete(&entity.BusinessHour{}).Error
}

func NewBusinessHourRepo(db *gorm.DB) BusinessHourRepositoryInterface {
	return &businessHourRepository{
		db: db,
	}
}
