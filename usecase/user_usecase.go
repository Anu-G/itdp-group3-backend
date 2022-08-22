package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type UserUsecase interface {
	Update(u *entity.User) error
	FindByUsername(u *entity.User) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: ur,
	}
}

func (uc *userUsecase) Update(u *entity.User) error {
	return uc.repo.Update(u)
}

func (uc *userUsecase) FindByUsername(u *entity.User) error {
	return uc.repo.FindAccountByUsername(u)
}
