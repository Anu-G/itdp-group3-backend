package entity

import "gorm.io/gorm"

type BussinessProfile struct {
	gorm.Model
	AccountID      uint   `gorm:"not null"`
	CategoryID     uint   `gorm:"not null"`
	Address        string `gorm:"size:250;not null"`
	ProfileImage   string `gorm:"size:250"`
	ProfileBio     string `gorm:"size:150"`
	GmapsLink      string `gorm:"size:250"`
	BussinessPhone string `gorm:"size:15"`

	BussinessHours []BussinessHour
	BussinessLinks []BussinessLink
	BussinessFAQs  []BussinessFAQ
}

func (bp BussinessProfile) TableName() string {
	return "m_bussiness_profile"
}
