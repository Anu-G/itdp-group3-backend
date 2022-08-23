package auth

import (
	"errors"
	"fmt"
	"itdp-group3-backend/config"
	"itdp-group3-backend/model/entity"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Token interface {
	CreateAccessToken(cred *entity.User) (*entity.TokenDetails, error)
	VerifyAccessToken(tokenStr string) (*entity.AccessDetail, error)
	StoreAccessToken(accessUuid string, tokenDetail *entity.TokenDetails) error
	FetchAccessToken(accessDetail *entity.AccessDetail) (string, error)
	OpenAccessToken(accessDetail *entity.AccessDetail) (string, error)
	DeleteAccessToken(accessDetail *entity.AccessDetail) (string, error)
}

type token struct {
	cfg config.TokenConfig
}

func (ts *token) CreateAccessToken(uc *entity.User) (*entity.TokenDetails, error) {
	newTokenDetail := new(entity.TokenDetails)
	now := time.Now().Local()
	end := now.Add(ts.cfg.AccessTokenLifeTime)
	newTokenDetail.AtExpires = end.Unix()
	newTokenDetail.AccessUuid = uuid.NewString()

	claims := AuthClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    ts.cfg.ApplicationName,
			IssuedAt:  now.Unix(),
			ExpiresAt: end.Unix(),
		},
		Username:   uc.Username,
		AccountID:  uint(uc.Account.ID),
		Email:      uc.Email,
		Role:       uint(uc.Account.RoleID),
		AccessUuid: newTokenDetail.AccessUuid,
	}

	token := jwt.NewWithClaims(ts.cfg.JwtSigningMethod, claims)
	if newToken, err := token.SignedString([]byte(ts.cfg.JwtSignatureKey)); err != nil {
		return nil, err
	} else {
		newTokenDetail.AccessToken = newToken
		return newTokenDetail, nil
	}
}

func (ts *token) VerifyAccessToken(tokenStr string) (*entity.AccessDetail, error) {
	newAccessDetail := new(entity.AccessDetail)
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		} else if method != ts.cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(ts.cfg.JwtSignatureKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != ts.cfg.ApplicationName {
		return nil, err
	}
	newAccessDetail.AccessUuid = claims["accessUUID"].(string)
	newAccessDetail.Role = uint(claims["role"].(float64))
	newAccessDetail.Username = claims["userName"].(string)

	return newAccessDetail, nil
}

func (ts *token) StoreAccessToken(accessUuid string, tokenDetail *entity.TokenDetails) error {
	now := time.Now().Local()
	end := time.Unix(tokenDetail.AtExpires, 0)
	if err := ts.cfg.Redis.Set(tokenDetail.AccessUuid, accessUuid, end.Sub(now)).Err(); err != nil {
		return err
	}
	return nil
}

func (ts *token) FetchAccessToken(accessDetail *entity.AccessDetail) (string, error) {
	if accessDetail != nil {
		userId, err := ts.cfg.Redis.Get(accessDetail.AccessUuid).Result()
		if err != nil {
			return "", err
		}
		return userId, nil
	} else {
		return "", fmt.Errorf("invalid access")
	}
}

func (ts *token) OpenAccessToken(accessDetail *entity.AccessDetail) (string, error) {
	if accessDetail != nil {
		userName, err := ts.cfg.Redis.Get(accessDetail.Username).Result()
		if err != nil {
			return "", err
		}
		return userName, nil
	} else {
		return "", fmt.Errorf("invalid access")
	}
}

func (ts *token) DeleteAccessToken(accessDetail *entity.AccessDetail) (string, error) {
	if accessDetail != nil {
		_, err := ts.cfg.Redis.Del(accessDetail.AccessUuid).Result()
		return accessDetail.AccessUuid, err
	} else {
		return accessDetail.AccessUuid, errors.New("no token can be found")
	}
}

func NewTokenService(cfg config.TokenConfig) Token {
	newToken := new(token)
	newToken.cfg = cfg
	return newToken
}
