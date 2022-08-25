package entity

import "gorm.io/gorm"

type Follower struct {
	gorm.Model
	AccountID         uint `gorm:"not null" json:"account_id"`
	FollowerAccountID uint `json:"follower_account_id"`
}

func (f Follower) TableName() string {
	return "m_follower"
}
