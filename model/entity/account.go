package entity

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	User        User   `gorm:"foreignkey:Username" json:"user_name"`
	RoleID      int    `gorm:"not null" json:"role_id"`
	DisplayName string `gorm:"size:36;not null" json:"display_name"`
	PhoneNumber string `gorm:"size:15;unique;not null" json:"phone_number"`

	BusinessProfile BusinessProfile
	Products        []Product
	Feeds           []Feed
}

func (a Account) TableName() string {
	return "m_account"
}
