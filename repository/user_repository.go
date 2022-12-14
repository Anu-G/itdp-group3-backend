package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Update(u *entity.User) error
	FindAccountByUsername(u *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Update(u *entity.User) error {
	return ur.db.Where("username = ?", u.Username).Updates(u).Error
}

func (ur *userRepository) FindAccountByUsername(u *entity.User) error {
	return ur.db.First(&u, "username = ?", u.Username).Error
}
