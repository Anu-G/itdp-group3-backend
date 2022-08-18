package entity

import (
	"encoding/base64"
	"log"
)

type UserCredential struct {
	ID       uint   `json:"id"`
	Username string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     uint   `json:"role"`
}

func (uc *UserCredential) Encode() {
	uc.Password = base64.RawStdEncoding.EncodeToString([]byte(uc.Password))
}

func (uc *UserCredential) Decode() {
	data, err := base64.StdEncoding.DecodeString(uc.Password)
	if err != nil {
		log.Println(err)
	}
	uc.Password = string(data)
}
