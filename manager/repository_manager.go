package manager

import "itdp-group3-backend/repository"

type RepositoryManagerInterface interface {
	BusinessProfileRepo() repository.BusinessProfileRepositoryInterface
	FileRepo() repository.FileRepository
}

type repositoryManager struct {
	dbCon InfraManagerInterface
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
