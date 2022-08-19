package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Update(a *entity.Account) error
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
	return ar.db.Updates(&a).Error
}
