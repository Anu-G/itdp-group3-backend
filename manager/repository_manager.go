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
}

type repositoryManager struct {
	dbCon InfraManagerInterface
}

func (rm *repositoryManager) ProductRepo() repository.ProductRepositoryInterface {
	return repository.NewProductRepo(rm.dbCon.DBCon())
}

func (rm *repositoryManager) BusinessProfileRepo() repository.BusinessProfileRepositoryInterface {
	return repository.NewBusinessProfileRepo(rm.dbCon.DBCon())
}

func (rm *repositoryManager) FileRepo() repository.FileRepository {
	return repository.NewFileRepository(`E:\ITDP Sinarmas Mining\toktok_dev\img`)
}

// NewRepo : init new repository manager
func NewRepo(dbCon InfraManagerInterface) RepositoryManagerInterface {
	return &repositoryManager{
		dbCon: dbCon,
	}
}

func (r *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.dbCon.DBCon())
}

func (r *repositoryManager) AuthRepo() repository.AuthRepository {
	return repository.NewAuthRepo(r.dbCon.DBCon())
}

func (r *repositoryManager) AccountRepo() repository.AccountRepository {
	return repository.NewAccountRepository(r.dbCon.DBCon())
}

func (r *repositoryManager) FeedRepo() repository.FeedRepository {
	return repository.NewFeedRepository(r.dbCon.DBCon())
}

func (r *repositoryManager) DetailMediaFeedRepo() repository.DetailMediaFeedRepository {
	return repository.NewDetailMediaFeedRepository(r.dbCon.DBCon())
}

func (r *repositoryManager) DetailCommentRepo() repository.DetailCommentRepository {
	return repository.NewDetailCommentRepository(r.dbCon.DBCon())
}
