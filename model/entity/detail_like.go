package entity

import "gorm.io/gorm"

type DetailLike struct {
	gorm.Model
	FeedID    uint `gorm:"not null" json:"feed_id"`
	AccountID uint `json:"account_id"`
}
