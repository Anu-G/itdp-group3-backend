package entity

import "gorm.io/gorm"

type BussinessLink struct {
	gorm.Model
	BussinessProfileID uint   `gorm:"not null"`
	label              string `gorm:"size:36;not null"`
	link               string `gorm:"size:250;not null"`
}

func (bl BussinessLink) TableName() string {
	return "m_bussiness_link"
}
