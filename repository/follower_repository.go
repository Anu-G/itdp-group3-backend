package repository

import (
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type FollowerRepository interface {
	Create(fl *entity.Follower) error
	Delete(fl *entity.Follower) error
	FindForVerif(fl *dto.FollowRequest) (entity.Follower, error)
}

type followerRepository struct {
	db *gorm.DB
}

func NewFollowerRepository(db *gorm.DB) FollowerRepository {
	return &followerRepository{
		db: db,
	}
}

func (flr *followerRepository) Create(fl *entity.Follower) error {
	return flr.db.Create(&fl).Error
}

func (flr *followerRepository) Delete(fl *entity.Follower) error {
	return flr.db.Where("account_id = ? AND follower_account_id = ?", fl.AccountID, fl.FollowerAccountID).Delete(&fl).Error
}

func (flr *followerRepository) FindForVerif(fl *dto.FollowRequest) (entity.Follower, error) {
	var followHold *entity.Follower
	res := flr.db.Where("account_id = ? AND follower_account_id = ?", fl.FollowedAccountID, fl.FollowerAccounID).Find(&followHold).Error
	return *followHold, res
}
