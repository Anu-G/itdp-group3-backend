package entity

import "gorm.io/gorm"

type Followed struct {
	gorm.Model
	AccountID         uint `gorm:"not null" json:"account_id"`
	FollowedAccountID uint `json:"following_account_id"`
}

func (f Followed) TableName() string {
	return "m_following"
}
