package auth

import "github.com/golang-jwt/jwt"

type AuthClaims struct {
	jwt.StandardClaims
	Username   string `json:"userName"`
	AccountID  uint   `json:"account_id"`
	Email      string `json:"email"`
	Role       uint   `json:"role"`
	AccessUuid string `json:"accessUUID"`
}
