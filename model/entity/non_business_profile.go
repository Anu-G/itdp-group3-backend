package entity

import "gorm.io/gorm"

type NonBusinessProfile struct {
	gorm.Model
	AccountID    uint   `gorm:"not null" json:"account_id"`
	ProfileImage string `gorm:"size:250" json:"profile_image"`
	ProfileBio   string `gorm:"size:150" json:"profile_bio"`
	DisplayName  string `gorm:"size:36;not null" json:"display_name"`
}

func (nbp NonBusinessProfile) TableName() string {
	return "m_non_business_profile"
}
