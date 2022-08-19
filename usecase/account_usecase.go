package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type AccountUsecase interface {
	Update(a *entity.Account) error
	FindByUsername(a *entity.Account) error
}

type accountUsecase struct {
	repo repository.AccountRepository
}

func NewAccountUsecse(repo repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		repo: repo,
	}
}

func (ac *accountUsecase) Update(a *entity.Account) error {
	return ac.repo.Update(a)
}

func (ac *accountUsecase) FindByUsername(a *entity.Account) error {
	return ac.repo.FindByUsername(a)
}
