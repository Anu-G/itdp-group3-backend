package entity

import "gorm.io/gorm"

type BusinessProfile struct {
	gorm.Model
	AccountID    uint   `gorm:"not null" json:"account_id"`
	CategoryID   uint   `gorm:"not null" json:"category_id"`
	Address      string `gorm:"size:250;not null" json:"address"`
	ProfileImage string `gorm:"size:250" json:"profile_image"`
	ProfileBio   string `gorm:"size:150" json:"profile_bio"`
	GmapsLink    string `gorm:"size:250" json:"gmaps_link"`
	DisplayName  string `gorm:"size:36;not null" json:"display_name"`

	BusinessHours []BusinessHour `json:"business_hours"`
	BusinessLinks []BusinessLink `json:"business_links"`
	BusinessFAQs  []BusinessFAQ  `json:"business_faqs"`
}

func (bp BusinessProfile) TableName() string {
	return "m_business_profile"
}
