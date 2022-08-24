package entity

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"itdp-group3-backend/utils"
	"log"

	"gorm.io/gorm"
)

var bytes = []byte(utils.CallGlobalVar().EncryptByte)

type User struct {
	gorm.Model
	Username string `gorm:"primaryKey;size:50;unique;not null" json:"user_name"`
	Password string `gorm:"size:50;not null" json:"password"`
	Email    string `gorm:"size:100;unique;not null" json:"email"`

	Account Account `gorm:"foreignkey:Username;references:Username" json:"account"`
}

func (uc User) TableName() string {
	return "m_user_credential"
}

func (uc *User) Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func (uc *User) Encrypt() {
	block, _ := aes.NewCipher([]byte(utils.SecretPassword))
	hiddenText := []byte(uc.Password)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(hiddenText))
	cfb.XORKeyStream(cipherText, hiddenText)
	res := uc.Encode(cipherText)
	uc.Password = res
}

func (uc *User) Decode(input string) []byte {
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Println(err)
	}
	return data
}

func (uc *User) Decrypt() {
	block, _ := aes.NewCipher([]byte(utils.SecretPassword))
	res := uc.Decode(uc.Password)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	cipherText := make([]byte, len(res))
	cfb.XORKeyStream(cipherText, res)
	uc.Password = string(cipherText)
}
