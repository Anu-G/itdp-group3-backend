package entity

import (
	"gorm.io/gorm"
)

type BusinessHour struct {
	gorm.Model
	BusinessProfileID uint   `gorm:"not null" json:"business_profile_id"`
	Day               int    `gorm:"not null" json:"day"`
	OpenHour          string `gorm:"not null" json:"open_hour"`
	CloseHour         string `gorm:"not null" json:"close_hour"`
}

func (bh BusinessHour) TableName() string {
	return "m_business_hours"
}
