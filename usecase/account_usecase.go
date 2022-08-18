package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type AccountUsecase interface {
	Create(a *entity.Account) error
}

type accountUsecase struct {
	repo repository.AccountRepository
}

func NewAccountUsecse(repo repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		repo: repo,
	}
}

func (ac *accountUsecase) Create(a *entity.Account) error {
	return ac.repo.Create(a)
}
