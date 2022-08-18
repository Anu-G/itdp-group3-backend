package entity

import "gorm.io/gorm"

type DetailMediaFeed struct {
	gorm.Model
	FeedID uint `gorm:"not null"`
	MediaLink string `gorm:"size:150;not null"`
}

func (dmf DetailMediaFeed) TableName() string {
	return "m_detail_media_feed"
}