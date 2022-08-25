package manager

import "itdp-group3-backend/usecase"

type UseCaseManagerInterface interface {
	UserUsecase() usecase.UserUsecase
	AuthUsecase() usecase.AuthUsecase
	AccountUsecase() usecase.AccountUsecase
	FeedUsecase() usecase.FeedUsecase
	DetailMediaFeedUsecase() usecase.DetailMediaFeedUsecase
	DetailCommentUsecase() usecase.DetailCommentUsecase
	BusinessProfileUseCase() usecase.BusinessProfileUseCaseInterface
	ProductUseCase() usecase.ProductUseCaseInterface
	NonBusinessProfileUseCase() usecase.NonBusinessProfileUseCaseInterface
	CategoryUsecase() usecase.CategoryUsecase
	FollowUsecase() usecase.FollowUsecase
}

type useCaseManager struct {
	repo RepositoryManagerInterface
}

func (um *useCaseManager) ProductUseCase() usecase.ProductUseCaseInterface {
	return usecase.NewProductUseCase(um.repo.ProductRepo(), um.repo.FileProductRepo())
}

func (um *useCaseManager) BusinessProfileUseCase() usecase.BusinessProfileUseCaseInterface {
	return usecase.NewBusinessProfileUseCase(um.repo.BusinessProfileRepo(), um.repo.AccountRepo(), um.repo.BusinessHourRepo(), um.repo.BusinessLinkRepo(), um.repo.FileRepo())
}

func (um *useCaseManager) NonBusinessProfileUseCase() usecase.NonBusinessProfileUseCaseInterface {
	return usecase.NewNonBusinessProfileUseCase(um.repo.NonBusinessProfileRepo(), um.repo.AccountRepo(), um.repo.FileRepo())
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

func (uc *useCaseManager) FeedUsecase() usecase.FeedUsecase {
	return usecase.NewFeedUsecase(uc.repo.FeedRepo(), uc.repo.AccountRepo())
}

func (uc useCaseManager) DetailMediaFeedUsecase() usecase.DetailMediaFeedUsecase {
	return usecase.NewDetailMediaFeedUsecase(uc.repo.DetailMediaFeedRepo(), uc.repo.FileRepo())
}

func (uc useCaseManager) DetailCommentUsecase() usecase.DetailCommentUsecase {
	return usecase.NewDetailCommentUsecase(uc.repo.DetailCommentRepo())
}

func (um useCaseManager) CategoryUsecase() usecase.CategoryUsecase {
	return usecase.NewCategoryUsecase(um.repo.CategoryRepo())
}

func (um *useCaseManager) FollowUsecase() usecase.FollowUsecase {
	return usecase.NewFollowUsecase(um.repo.FollowerRepo(), um.repo.FollowedRepo())
}
