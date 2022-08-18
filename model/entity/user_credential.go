package entity

import (
	"encoding/base64"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"primaryKey;size:50;unique;not null"`
	Password string `gorm:"size:50;not null"`
	Email    string `gorm:"size:100;unique;not null"`

	Account Account `gorm:"foreignkey:Username;references:Username"`
}

func (uc User) TableName() string {
	return "m_user_credential"
}

func (uc *User) Encode() {
	uc.Password = base64.StdEncoding.EncodeToString([]byte(uc.Password))
}

func (uc *User) Decode() {
	data, err := base64.StdEncoding.DecodeString(uc.Password)
	if err != nil {
		log.Println(err)
	}

	uc.Password = string(data)
}
