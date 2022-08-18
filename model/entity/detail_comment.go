package entity

import "gorm.io/gorm"

type DetailComment struct {
	gorm.Model
	FeedID uint `gorm:"not null" json:"feed_id"`
	CommentFill string `gorm:"size:500;not null" json:"comment_fill"`
}

func (dc DetailComment) TableName() string {
	return "m_detail_comment"
}