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
	FileProductRepo() repository.FileProductRepository
	NonBusinessProfileRepo() repository.NonBusinessProfileRepositoryInterface
	CategoryRepo() repository.CategoryRepository
	BusinessHourRepo() repository.BusinessHourRepositoryInterface
	BusinessLinkRepo() repository.BusinessLinkRepositoryInterface
}

type repositoryManager struct {
	infra InfraManagerInterface
}

func (rm *repositoryManager) BusinessHourRepo() repository.BusinessHourRepositoryInterface {
	return repository.NewBusinessHourRepo(rm.infra.DBCon())
}

func (rm *repositoryManager) BusinessLinkRepo() repository.BusinessLinkRepositoryInterface {
	return repository.NewBusinessLinkRepo(rm.infra.DBCon())
}

func (rm *repositoryManager) ProductRepo() repository.ProductRepositoryInterface {
	return repository.NewProductRepo(rm.infra.DBCon())
}

func (rm *repositoryManager) BusinessProfileRepo() repository.BusinessProfileRepositoryInterface {
	return repository.NewBusinessProfileRepo(rm.infra.DBCon())
}

func (rm *repositoryManager) FileProductRepo() repository.FileProductRepository {
	return repository.NewFileProductRepository(rm.infra.GetMediaPathProduct())
}

func (rm *repositoryManager) FileRepo() repository.FileRepository {
	return repository.NewFileRepository(rm.infra.GetMediaPath(), rm.infra.GetMediaPathFeed())
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

func (r *repositoryManager) CategoryRepo() repository.CategoryRepository {
	return repository.NewCategoryRepository(r.infra.DBCon())
}
