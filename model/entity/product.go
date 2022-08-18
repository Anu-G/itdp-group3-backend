package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	AccountID   uint    `gorm:"not null"`
	ProductName string  `gorm:"size:100;not null"`
	Price       float64 `gorm:"0;not null"`
	Description string  `gorm:"size:256;not null"`

	DetailMediaProducts []DetailMediaProduct
}

func (p Product) TableName() string {
	return "m_product"
}