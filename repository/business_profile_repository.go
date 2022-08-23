package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type BusinessProfileRepositoryInterface interface {
	Create(bp *entity.BusinessProfile) error
	GetByIdPreload(bp *entity.BusinessProfile) error
	Delete(id string) error
}

type businessProfileRepository struct {
	db *gorm.DB
}

func (b *businessProfileRepository) Delete(id string) error {
	return b.db.Unscoped().Where("account_id = ?", id).Delete(&entity.BusinessProfile{}).Error
}

func (b *businessProfileRepository) GetByIdPreload(bp *entity.BusinessProfile) error {
	return b.db.Preload("BusinessHours").Preload("BusinessLinks").Preload("BusinessFAQs").First(&bp, "m_business_profile.account_id = ?", bp.AccountID).Error
}

func (b *businessProfileRepository) Create(bp *entity.BusinessProfile) error {
	return b.db.Create(&bp).Error
}

func NewBusinessProfileRepo(db *gorm.DB) BusinessProfileRepositoryInterface {
	return &businessProfileRepository{
		db: db,
	}
}
