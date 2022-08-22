package manager

import "itdp-group3-backend/usecase"

type UseCaseManagerInterface interface {
	UserUsecase() usecase.UserUsecase
	AuthUsecase() usecase.AuthUsecase
	AccountUsecase() usecase.AccountUsecase
	FeedUsecase() usecase.FeedUsecase
	DetailMediaFeedUsecase() usecase.DetailMediaFeedUsecase
	DetailCommentUsecase() usecase.DetailCommentUsecase
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

func (uc *useCaseManager) FeedUsecase() usecase.FeedUsecase {
	return usecase.NewFeedUsecase(uc.repo.FeedRepo())
}

func (uc useCaseManager) DetailMediaFeedUsecase() usecase.DetailMediaFeedUsecase {
	return usecase.NewDetailMediaFeedUsecase(uc.repo.DetailMediaFeedRepo())
}

func (uc useCaseManager) DetailCommentUsecase() usecase.DetailCommentUsecase {
	return usecase.NewDetailCommentUsecase(uc.repo.DetailCommentRepo())
}
