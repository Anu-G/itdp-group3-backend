package usecase

import (
	"itdp-group3-backend/repository"
)

type UserUsecase interface {
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: ur,
	}
}
