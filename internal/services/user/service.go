package user

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/repositories"
	repouser "github.com/DenisKhanov/PrivateKeeper/internal/repositories/database/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

// checking interface compliance at the compiler level
var _ repositories.UserRepository = (*repouser.RepositoryUser)(nil)

// ServiceUser ...
type ServiceUser struct {
	repository repositories.UserRepository
	dbPool     *pgxpool.Pool
}

// NewServiceUser .....
func NewServiceUser(repository repositories.UserRepository) *ServiceUser {
	return &ServiceUser{
		repository: repository,
	}
}
