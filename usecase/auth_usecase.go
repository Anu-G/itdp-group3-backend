package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type AuthUsecase interface {
	CreateUser(uc *entity.UserCredential) error
	FindUser(uc *entity.UserCredential) (entity.UserCredential, error)
}

type authUsecase struct {
	repo repository.AuthRepository
}

func (au *authUsecase) CreateUser(uc *entity.UserCredential) error {
	return au.repo.CreateUser(uc)
}

func (au *authUsecase) FindUser(uc *entity.UserCredential) (entity.UserCredential, error) {
	return au.repo.FindUser(uc)
}
