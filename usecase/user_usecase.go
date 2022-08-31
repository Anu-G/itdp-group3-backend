package usecase

import (
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type UserUsecase interface {
	Update(u *dto.UpdateUserRequest) error
	FindByUsername(u *entity.User) error
}

type userUsecase struct {
	repo    repository.UserRepository
	repoAcc repository.AccountRepository
}

func NewUserUsecase(ur repository.UserRepository, urA repository.AccountRepository) UserUsecase {
	return &userUsecase{
		repo:    ur,
		repoAcc: urA,
	}
}

func (uc *userUsecase) Update(u *dto.UpdateUserRequest) error {
	var newUser entity.User
	var newAccount entity.Account
	newAccount.ID = u.AccountID
	err := uc.repoAcc.FindById(&newAccount)
	if err != nil {
		return err
	}
	newUser.Password = u.Password
	newUser.Username = u.Username
	newUser.Email = u.Email
	newAccount.PhoneNumber = u.PhoneNumber
	newUser.Encrypt()
	err = uc.repo.Update(&newUser)
	if err != nil {
		return err
	}
	err = uc.repoAcc.Update(&newAccount)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUsecase) FindByUsername(u *entity.User) error {
	return uc.repo.FindAccountByUsername(u)
}
