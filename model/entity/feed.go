package entity

import "gorm.io/gorm"

type Feed struct {
	gorm.Model
	AccountID   uint   `gorm:"not null"`
	CaptionPost string `gorm:"size:280;not null"`

	DetailMediaFeeds []DetailMediaFeed
	DetailComments   []DetailComment
}

func (f Feed) TableName() string {
	return "m_feed"
}
