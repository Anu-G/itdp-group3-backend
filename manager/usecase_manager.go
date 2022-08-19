package manager

import "itdp-group3-backend/usecase"

type UseCaseManagerInterface interface {
	UserUsecase() usecase.UserUsecase
	AuthUsecase() usecase.AuthUsecase
	AccountUsecase() usecase.AccountUsecase
}

type useCaseManager struct {
	repo RepositoryManagerInterface
}

// NewUseCase : init new use case manager
func NewUseCase(manager RepositoryManagerInterface) UseCaseManagerInterface {
	return &useCaseManager{
		repo: manager,
	}
}

func (uc *useCaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(uc.repo.UserRepo())
}

func (uc *useCaseManager) AuthUsecase() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(uc.repo.AuthRepo())
}

func (uc *useCaseManager) AccountUsecase() usecase.AccountUsecase {
	return usecase.NewAccountUsecse(uc.repo.AccountRepo())
}
