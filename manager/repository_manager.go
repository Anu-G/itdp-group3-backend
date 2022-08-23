package manager

import "itdp-group3-backend/repository"

type RepositoryManagerInterface interface {
	UserRepo() repository.UserRepository
	AuthRepo() repository.AuthRepository
	AccountRepo() repository.AccountRepository
	FeedRepo() repository.FeedRepository
	DetailMediaFeedRepo() repository.DetailMediaFeedRepository
	DetailCommentRepo() repository.DetailCommentRepository
	BusinessProfileRepo() repository.BusinessProfileRepositoryInterface
	ProductRepo() repository.ProductRepositoryInterface
	FileRepo() repository.FileRepository
	NonBusinessProfileRepo() repository.NonBusinessProfileRepositoryInterface
}

type repositoryManager struct {
	infra InfraManagerInterface
}

func (rm *repositoryManager) ProductRepo() repository.ProductRepositoryInterface {
	return repository.NewProductRepo(rm.infra.DBCon())
}

func (rm *repositoryManager) BusinessProfileRepo() repository.BusinessProfileRepositoryInterface {
	return repository.NewBusinessProfileRepo(rm.infra.DBCon())
}

func (rm *repositoryManager) FileRepo() repository.FileRepository {
	return repository.NewFileRepository(rm.infra.GetMediaPath())
}

func (rm *repositoryManager) NonBusinessProfileRepo() repository.NonBusinessProfileRepositoryInterface {
	return repository.NewNonBusinessProfileRepo(rm.infra.DBCon())
}

// NewRepo : init new repository manager
func NewRepo(infra InfraManagerInterface) RepositoryManagerInterface {
	return &repositoryManager{
		infra: infra,
	}
}

func (r *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.DBCon())
}

func (r *repositoryManager) AuthRepo() repository.AuthRepository {
	return repository.NewAuthRepo(r.infra.DBCon())
}

func (r *repositoryManager) AccountRepo() repository.AccountRepository {
	return repository.NewAccountRepository(r.infra.DBCon())
}

func (r *repositoryManager) FeedRepo() repository.FeedRepository {
	return repository.NewFeedRepository(r.infra.DBCon())
}

func (r *repositoryManager) DetailMediaFeedRepo() repository.DetailMediaFeedRepository {
	return repository.NewDetailMediaFeedRepository(r.infra.DBCon())
}

func (r *repositoryManager) DetailCommentRepo() repository.DetailCommentRepository {
	return repository.NewDetailCommentRepository(r.infra.DBCon())
}
