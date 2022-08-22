package entity

import "gorm.io/gorm"

type DetailMediaFeed struct {
	gorm.Model
	FeedID    uint   `gorm:"not null" json:"feed_id"`
	MediaLink string `gorm:"size:150;not null" json:"media_link"`
}

func (dmf DetailMediaFeed) TableName() string {
	return "m_detail_media_feed"
}
