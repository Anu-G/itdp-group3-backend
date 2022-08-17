package manager

type UseCaseManagerInterface interface {
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
