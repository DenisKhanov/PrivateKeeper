package data

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/repositories"
	repodata "github.com/DenisKhanov/PrivateKeeper/internal/repositories/database/data"
	repouser "github.com/DenisKhanov/PrivateKeeper/internal/repositories/database/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

// checking interface compliance at the compiler level
var _ repositories.UserRepository = (*repouser.RepositoryUser)(nil)
var _ repositories.DataRepository = (*repodata.RepositoryData)(nil)

// ServiceData ...
type ServiceData struct {
	repository   repositories.DataRepository
	s3Repository repositories.S3Repository
	dbPool       *pgxpool.Pool
}

// NewServiceData .....
func NewServiceData(repository repositories.DataRepository, s3 repositories.S3Repository, dbPool *pgxpool.Pool) *ServiceData {
	return &ServiceData{
		repository:   repository,
		s3Repository: s3,
		dbPool:       dbPool,
	}
}
