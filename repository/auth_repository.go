package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(uc *entity.User) error
	FindUser(uc *entity.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (ar *authRepository) CreateUser(u *entity.User) error {
	trx := ar.db.Begin()
	if err := ar.db.Create(&u).Error; err != nil {
		trx.Callback()
		return err
	}
	if err := ar.db.Create(&u.Account).Error; err != nil {
		trx.Callback()
		return err
	}
	trx.Commit()
	return nil
}

func (ar *authRepository) FindUser(u *entity.User) error {
	u.Encode()
	res := ar.db.Preload("Account").First(&u, "username = ?", u.Username)
	return res.Error
}
