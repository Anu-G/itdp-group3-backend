package entity

import "gorm.io/gorm"

type BusinessLink struct {
	gorm.Model
	BusinessProfileID uint   `gorm:"not null" json:"business_profile_id"`
	Label             string `gorm:"size:36;not null" json:"label"`
	Link              string `gorm:"size:250;not null" json:"link"`
}

func (bl BusinessLink) TableName() string {
	return "m_business_link"
}
