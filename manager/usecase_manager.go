package manager

import "itdp-group3-backend/usecase"

type UseCaseManagerInterface interface {
	BusinessProfileUseCase() usecase.BusinessProfileUseCaseInterface
	NonBusinessProfileUseCase() usecase.NonBusinessProfileUseCaseInterface
	ProductUseCase() usecase.ProductUseCaseInterface
}

type useCaseManager struct {
	repo RepositoryManagerInterface
}

func (um *useCaseManager) ProductUseCase() usecase.ProductUseCaseInterface {
	return usecase.NewProductUseCase(um.repo.ProductRepo())
}

func (um *useCaseManager) NonBusinessProfileUseCase() usecase.NonBusinessProfileUseCaseInterface {
	return usecase.NewNonBusinessProfileUseCase(um.repo.NonBusinessProfileRepo(), um.repo.FileRepo())
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
