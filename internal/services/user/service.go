package user

import (
	"crypto/rsa"
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
	publicKey  *rsa.PublicKey
}

// NewServiceUser .....
func NewServiceUser(repository repositories.UserRepository, publicKey *rsa.PublicKey) *ServiceUser {
	return &ServiceUser{
		repository: repository,
		publicKey:  publicKey,
	}
}
