package entity

import "gorm.io/gorm"

type DetailMediaProduct struct {
	gorm.Model
	ProductID uint `gorm:"not null"`
	MediaLink string `gorm:"size:150;not null"`
}

func (dmp DetailMediaProduct) TableName() string {
	return "m_detail_media_product"
}