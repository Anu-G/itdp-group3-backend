package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type BusinessLinkRepositoryInterface interface {
	Delete(id string) error
}

type businessLinkRepository struct {
	db *gorm.DB
}

func (b *businessLinkRepository) Delete(id string) error {
	return b.db.Unscoped().Where("id = ?", id).Delete(&entity.BusinessLink{}).Error
}

func NewBusinessLinkRepo(db *gorm.DB) BusinessLinkRepositoryInterface {
	return &businessLinkRepository{
		db: db,
	}
}
