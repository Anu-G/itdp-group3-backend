package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type BusinessProfileRepositoryInterface interface {
	Create(bp *entity.BusinessProfile) error
}

type businessProfileRepository struct {
	db *gorm.DB
}

func (b *businessProfileRepository) Create(bp *entity.BusinessProfile) error {
	return b.db.Create(&bp).Error
}

func NewBusinessProfileRepo(db *gorm.DB) BusinessProfileRepositoryInterface {
	return &businessProfileRepository{
		db: db,
	}
}
