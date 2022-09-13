package repository

import (
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type DetailLikeRepository interface {
	Like(dl *entity.DetailLike) error
	Unlike(dl *entity.DetailLike) error
	Find(dl *dto.LikeRequest) (entity.DetailLike, error)
}

type detailLikeRepository struct {
	db *gorm.DB
}

func NewDetailLikeRepository(db *gorm.DB) DetailLikeRepository {
	return &detailLikeRepository{
		db: db,
	}
}

func (dlr *detailLikeRepository) Like(dl *entity.DetailLike) error {
	return dlr.db.Create(&dl).Error
}

func (dlr *detailLikeRepository) Unlike(dl *entity.DetailLike) error {
	return dlr.db.Where("feed_id = ? AND account_id = ?", dl.FeedID, dl.AccountID).Delete(&dl).Error
}

func (dlr *detailLikeRepository) Find(dl *dto.LikeRequest) (entity.DetailLike, error) {
	var like *entity.DetailLike
	res := dlr.db.Where("feed_id = ? AND account_id = ?", dl.FeedID, dl.AccountID).First(&like)
	return *like, res.Error
}
