package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(uc *entity.UserCredential) error
	FindUser(uc *entity.UserCredential) (entity.UserCredential, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (ar *authRepository) CreateUser(uc *entity.UserCredential) error {
	return ar.db.Create(&uc).Error
}

func (ar *authRepository) FindUser(uc *entity.UserCredential) (entity.UserCredential, error) {
	uc.Encode()
	res := ar.db.First(&uc, "user_name = ?", uc.Username)
	return *uc, res.Error
}
