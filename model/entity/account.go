package entity

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username    string `gorm:"not null" json:"user_name"`
	RoleID      int    `gorm:"not null" json:"role_id"`
	PhoneNumber string `gorm:"size:15;unique;not null" json:"phone_number"`

	BusinessProfile    BusinessProfile    `json:"business_profile"`
	NonBusinessProfile NonBusinessProfile `json:"non_business_profile"`
	Products           []Product          `json:"products"`
	Feeds              []Feed             `json:"feeds"`
}

func (a Account) TableName() string {
	return "m_account"
}
