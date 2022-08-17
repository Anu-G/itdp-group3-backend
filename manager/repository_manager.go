package manager

type RepositoryManagerInterface interface {
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
