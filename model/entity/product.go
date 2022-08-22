package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	AccountID   uint    `gorm:"not null" json:"account_id"`
	ProductName string  `gorm:"size:100;not null" json:"product_name"`
	Price       float64 `gorm:"0;not null" json:"price"`
	Description string  `gorm:"size:256;not null" json:"description"`
	DetailMediaProducts string `json:"detail_media_products"`
}

func (p Product) TableName() string {
	return "m_product"
}