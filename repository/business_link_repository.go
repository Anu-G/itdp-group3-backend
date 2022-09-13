package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type BusinessLinkRepositoryInterface interface {
	Create(links *entity.BusinessLink) error
	Delete(id string) error
}

type businessLinkRepository struct {
	db *gorm.DB
}

func (b *businessLinkRepository) Create(links *entity.BusinessLink) error {
	return b.db.Create(&links).Error
}

func (b *businessLinkRepository) Delete(id string) error {
	return b.db.Where("id = ?", id).Delete(&entity.BusinessLink{}).Error
}

func NewBusinessLinkRepo(db *gorm.DB) BusinessLinkRepositoryInterface {
	return &businessLinkRepository{
		db: db,
	}
}
