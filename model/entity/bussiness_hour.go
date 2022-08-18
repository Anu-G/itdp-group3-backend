package entity

import (
	"time"

	"gorm.io/gorm"
)

type BussinessHour struct {
	gorm.Model
	BussinessProfileID uint      `gorm:"not null"`
	Day                int       `gorm:"not null"`
	OpenHour           time.Time `gorm:"not null"`
	CloseHour          time.Time `gorm:"not null"`
}

func (bh BussinessHour) TableName() string {
	return "m_bussiness_hours"
}
