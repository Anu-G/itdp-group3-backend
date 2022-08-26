package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type FollowedRepository interface {
	Create(fl *entity.Followed) error
	Delete(fl *entity.Followed) error
}

type followedRepository struct {
	db *gorm.DB
}

func NewFollowedRepository(db *gorm.DB) FollowedRepository {
	return &followedRepository{
		db: db,
	}
}

func (flr *followedRepository) Create(fl *entity.Followed) error {
	return flr.db.Create(&fl).Error
}

func (flr *followedRepository) Delete(fl *entity.Followed) error {
	return flr.db.Where("account_id = ? AND followed_account_id = ?", fl.AccountID, fl.FollowedAccountID).Delete(&fl).Error
}
