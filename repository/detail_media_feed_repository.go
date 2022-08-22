package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type DetailMediaFeedRepository interface {
	Create(*entity.DetailMediaFeed) error
	Read(*entity.DetailMediaFeed) error
}

type detailMediaFeedRepository struct {
	db *gorm.DB
}

func NewDetailMediaFeedRepository(db *gorm.DB) DetailMediaFeedRepository {
	return &detailMediaFeedRepository{
		db: db,
	}
}

func (fmr *detailMediaFeedRepository) Create(fm *entity.DetailMediaFeed) error {
	return fmr.db.Create(&fm).Error
}

func (fmr *detailMediaFeedRepository) Read(fm *entity.DetailMediaFeed) error {
	return fmr.db.Find(&fm, "id = ?", fm.ID).Error
}
