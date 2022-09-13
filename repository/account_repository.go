package repository

import (
	"errors"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Update(a *entity.Account) error
	ReadForPostTimeline(a *[]entity.Account) error
	ReadForProductDetail(a *entity.Account) error
	ReadForFeedDetail(a *entity.Account) error
	FindByUsername(a *entity.Account) error
	FindById(a *entity.Account) error
	FindListById(ids []uint) ([]entity.Account, error)
	GetAccount(id uint) (dto.GetAccountResponse, error)
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

func (ar *accountRepository) ReadForPostTimeline(a *[]entity.Account) error {
	return ar.db.Preload("Follower").Preload("Followed").Preload("BusinessProfile").Preload("NonBusinessProfile").Preload("Feeds").Find(&a).Error
}

func (ar *accountRepository) ReadForProductDetail(a *entity.Account) error {
	return ar.db.Preload("Follower").Preload("Followed").Preload("BusinessProfile").Preload("NonBusinessProfile").Preload("Products").First(&a, "id = ?", a.ID).Error
}

func (ar *accountRepository) ReadForFeedDetail(a *entity.Account) error {
	return ar.db.Preload("Follower").Preload("Followed").Preload("BusinessProfile").Preload("NonBusinessProfile").Preload("Feeds").First(&a, "id = ?", a.ID).Error
}

func (ar *accountRepository) FindByUsername(a *entity.Account) error {
	return ar.db.Preload("Follower").Preload("Followed").First(&a, "username = ?", a.Username).Error
}

func (ar *accountRepository) FindById(a *entity.Account) error {
	return ar.db.Preload("Follower").Preload("Followed").First(&a, "id = ?", a.ID).Error
}

func (ar *accountRepository) FindListById(ids []uint) ([]entity.Account, error) {
	var accountList *[]entity.Account
	if len(ids) == 0 {
		return nil, nil
	}
	res := ar.db.Where("id = ?", ids[0])
	for i := 1; i < len(ids); i++ {
		res = res.Or("id = ?", ids[i])
	}
	res = res.Find(&accountList)
	return *accountList, res.Error
}

func (ar *accountRepository) GetAccount(id uint) (dto.GetAccountResponse, error) {
	var account dto.GetAccountResponse

	res := ar.db.Raw(`SELECT email, phone_number from m_user_credential a INNER JOIN m_account b ON a.username = b.username WHERE b.id = ?`, id).Scan(&account)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return account, nil
		} else {
			return account, err
		}
	}
	return account, nil
}
