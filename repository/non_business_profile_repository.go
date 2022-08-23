package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type NonBusinessProfileRepositoryInterface interface {
	Create(bp *entity.NonBusinessProfile) error
	GetById(bp *entity.NonBusinessProfile) error
	Update(bp *entity.NonBusinessProfile) error
	Delete(id string) error
}

type nonBusinessProfileRepository struct {
	db *gorm.DB
}

func (n *nonBusinessProfileRepository) Delete(id string) error {
	return n.db.Unscoped().Where("account_id = ?", id).Delete(&entity.NonBusinessProfile{}).Error
}

func (n *nonBusinessProfileRepository) Update(bp *entity.NonBusinessProfile) error {
	return n.db.Where("account_id = ?", bp.AccountID).Updates(bp).Error
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
