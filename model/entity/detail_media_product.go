package entity

import "gorm.io/gorm"

type DetailMediaProduct struct {
	gorm.Model
	ProductID uint `gorm:"not null" json:"product_id"`
	MediaLink string `gorm:"size:150;not null" json:"media_link"`
}

func (dmp DetailMediaProduct) TableName() string {
	return "m_detail_media_product"
}