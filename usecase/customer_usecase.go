package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type UserUsecase interface {
	CreateUser(u *entity.UserCredential) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: ur,
	}
}

func (uc userUsecase) CreateUser(u *entity.UserCredential) error {
	return uc.repo.Create(*u)
}
