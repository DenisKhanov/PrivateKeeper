package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// RepositoryUser ...
type RepositoryUser struct {
	dbPool *pgxpool.Pool //opened in main func dbPool pool connections
}

func NewPostgresUser(dbPool *pgxpool.Pool) (*RepositoryUser, error) {
	storage := &RepositoryUser{
		dbPool: dbPool,
	}
	return storage, nil
}
