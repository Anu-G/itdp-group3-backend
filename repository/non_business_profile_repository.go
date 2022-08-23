package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type NonBusinessProfileRepositoryInterface interface {
	Create(bp *entity.NonBusinessProfile) error
	GetById(bp *entity.NonBusinessProfile) error
	GetPhoneNumber(accountId uint) (acc entity.Account, err error)
}

type nonBusinessProfileRepository struct {
	db *gorm.DB
}

func (n *nonBusinessProfileRepository) GetPhoneNumber(accountId uint) (acc entity.Account, err error) {
	err = n.db.First(&acc, "m_account.id = ?", accountId).Error
	return acc, err
}

func (n *nonBusinessProfileRepository) GetById(bp *entity.NonBusinessProfile) error {
	return n.db.First(&bp, "m_non_business_profile.account_id = ?", bp.AccountID).Error
}

func (n *nonBusinessProfileRepository) Create(bp *entity.NonBusinessProfile) error {
	return n.db.Create(&bp).Error
}

func NewNonBusinessProfileRepo(db *gorm.DB) NonBusinessProfileRepositoryInterface {
	return &nonBusinessProfileRepository{
		db: db,
	}
}
