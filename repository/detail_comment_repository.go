package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type DetailCommentRepository interface {
	Create(cm *entity.DetailComment) error
	Read(cm *entity.DetailComment) error
}

type detailCommentRepository struct {
	db *gorm.DB
}

func NewDetailCommentRepository(db *gorm.DB) DetailCommentRepository {
	return &detailCommentRepository{
		db: db,
	}
}

func (cmr *detailCommentRepository) Create(cm *entity.DetailComment) error {
	return cmr.db.Create(&cm).Error
}

func (cmr *detailCommentRepository) Read(cm *entity.DetailComment) error {
	return cmr.db.Find(&cm, "id = ?", cm.ID).Error
}
