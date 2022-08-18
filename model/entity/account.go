package entity

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username    uint   `gorm:"not null"`
	RoleID      int    `gorm:"not null"`
	DisplayName string `gorm:"size:36;not null"`
	PhoneNumber string `gorm:"size:15;unique;not null"`

	BussinessProfile BussinessProfile
	Products         []Product
	Feeds            []Feed
}

func (a Account) TableName() string {
	return "m_account"
}
