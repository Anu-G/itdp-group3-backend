package entity

import "gorm.io/gorm"

type DetailComment struct {
	gorm.Model
	FeedID uint `gorm:"not null"`
	CommentFill string `gorm:"size:500;not null"`
}

func (dc DetailComment) TableName() string {
	return "m_detail_comment"
}