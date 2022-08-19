package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type UserUsecase interface {
	CreateUser(u *entity.User) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: ur,
	}
}

func (uc userUsecase) CreateUser(u *entity.User) error {
	return uc.repo.Create(u)
}
