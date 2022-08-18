package entity

import "gorm.io/gorm"

type BussinessFAQ struct {
	gorm.Model
	BussinessProfileID uint   `gorm:"not null"`
	Question           string `gorm:"size:250;not null"`
	Answer             string `gorm:"size:500;not null"`
}

func (bf BussinessFAQ) TableName() string {
	return "m_bussiness_faq"
}
