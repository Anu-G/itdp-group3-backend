package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `gorm:"size:100;not null" json:"category_name"`

	BusinessProfiles []BusinessProfile `json:"business_profiles"`
}

func (c Category) TableName() string {
	return "m_category"
}
