package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Update(a *entity.Account) error
	FindByUsername(a *entity.Account) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (ar *accountRepository) Update(a *entity.Account) error {
	return ar.db.Where("username = ?", a.Username).Updates(&a).Error
}

func (ar *accountRepository) FindByUsername(a *entity.Account) error {
	return ar.db.First(&a, "username = ?", a.Username).Error
}
