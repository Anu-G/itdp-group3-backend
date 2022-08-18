package manager

import "itdp-group3-backend/repository"

type RepositoryManagerInterface interface {
	UserRepo() repository.UserRepository
	AuthRepo() repository.AuthRepository
}

type repositoryManager struct {
	dbCon InfraManagerInterface
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
