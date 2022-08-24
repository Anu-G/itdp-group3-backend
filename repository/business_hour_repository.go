package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type BusinessHourRepositoryInterface interface {
	Delete(id string) error
}

type businessHourRepository struct {
	db *gorm.DB
}

func (b *businessHourRepository) Delete(id string) error {
	return b.db.Unscoped().Where("id = ?", id).Delete(&entity.BusinessHour{}).Error
}

func NewBusinessHourRepo(db *gorm.DB) BusinessHourRepositoryInterface {
	return &businessHourRepository{
		db: db,
	}
}
