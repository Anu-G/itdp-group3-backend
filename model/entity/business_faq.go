package entity

import "gorm.io/gorm"

type BusinessFAQ struct {
	gorm.Model
	BusinessProfileID uint   `gorm:"not null" json:"business_profile_id"`
	Question           string `gorm:"size:250;not null" json:"question"`
	Answer             string `gorm:"size:500;not null" json:"answer"`
}

func (bf BusinessFAQ) TableName() string {
	return "m_business_faq"
}
