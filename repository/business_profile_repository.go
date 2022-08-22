package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type BusinessProfileRepositoryInterface interface {
	Create(bp *entity.BusinessProfile) error
	GetByIdPreload(bp *entity.BusinessProfile) error
	GetPhoneNumber(accountId uint) (acc entity.Account, err error)
}

type businessProfileRepository struct {
	db *gorm.DB
}

func (b *businessProfileRepository) GetPhoneNumber(accountId uint) (acc entity.Account, err error) {
	err = b.db.First(&acc, "m_account.id = ?", accountId).Error
	return acc, err
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
