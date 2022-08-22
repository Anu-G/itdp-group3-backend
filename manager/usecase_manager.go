package manager

import "itdp-group3-backend/usecase"

type UseCaseManagerInterface interface {
	BusinessProfileUseCase() usecase.BusinessProfileUseCaseInterface
}

type useCaseManager struct {
	repo RepositoryManagerInterface
}

func (um *useCaseManager) BusinessProfileUseCase() usecase.BusinessProfileUseCaseInterface {
	return usecase.NewBusinessProfileUseCase(um.repo.BusinessProfileRepo(), um.repo.FileRepo())
}

// NewUseCase : init new use case manager
func NewUseCase(manager RepositoryManagerInterface) UseCaseManagerInterface {
	return &useCaseManager{
		repo: manager,
	}
}
