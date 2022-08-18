package auth

import "github.com/golang-jwt/jwt"

type AuthClaims struct {
	jwt.StandardClaims
	Username   string `json:"userName"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Role       uint   `json:"role"`
	AccessUuid string `json:"accessUUID"`
}
