package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `gorm:"size:100;not null"`

	BussinessProfiles []BussinessProfile
}

func (c Category) TableName() string {
	return "m_category"
}
