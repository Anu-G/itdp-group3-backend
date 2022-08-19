package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type AuthUsecase interface {
	CreateUser(uc *entity.User) error
	FindUser(uc *entity.User) error
}

type authUsecase struct {
	repo repository.AuthRepository
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &authUsecase{
		repo: repo,
	}
}

func (au *authUsecase) CreateUser(uc *entity.User) error {
	return au.repo.CreateUser(uc)
}

func (au *authUsecase) FindUser(uc *entity.User) error {
	return au.repo.FindUser(uc)
}
