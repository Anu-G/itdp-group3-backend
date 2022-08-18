package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(u entity.UserCredential) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Create(u entity.UserCredential) error {
	return ur.db.Create(u).Error
}
