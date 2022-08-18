package repository

import (
	"fmt"
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(a *entity.Account) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (ar *accountRepository) Create(a *entity.Account) error {
	fmt.Println("tes4")
	return ar.db.Create(&a).Error
}
