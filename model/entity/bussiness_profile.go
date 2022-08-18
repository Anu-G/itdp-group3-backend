package entity

import "gorm.io/gorm"

type BusinessProfile struct {
	gorm.Model
	AccountID      uint   `gorm:"not null" json:"account_id"`
	CategoryID     uint   `gorm:"not null" json:"category_id"`
	Address        string `gorm:"size:250;not null" json:"address"`
	ProfileImage   string `gorm:"size:250" json:"profile_image"`
	ProfileBio     string `gorm:"size:150" json:"profile_bio"`
	GmapsLink      string `gorm:"size:250" json:"gmaps_link"`
	BusinessPhone string `gorm:"size:15" json:"business_phone"`

	BusinessHours []BusinessHour
	BusinessLinks []BusinessLink
	BusinessFAQs  []BusinessFAQ
}

func (bp BusinessProfile) TableName() string {
	return "m_business_profile"
}
