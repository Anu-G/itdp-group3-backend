package entity

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
	AccountID   uint   `gorm:"not null" json:"account_id"`
	CaptionPost string `gorm:"size:280;not null" json:"caption_post"`

	DetailMediaFeeds []DetailMediaFeed `json:"detail_media_feeds"`
	DetailComments   []DetailComment   `json:"detail_comments"`
}

func (f Feed) TableName() string {
	return "m_feed"
}
