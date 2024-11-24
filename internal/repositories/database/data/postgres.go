package data

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// RepositoryData ...
type RepositoryData struct {
	dbPool *pgxpool.Pool //opened in main func dbPool pool connections
}

func NewPostgresData(dbPool *pgxpool.Pool) *RepositoryData {
	return &RepositoryData{
		dbPool: dbPool,
	}
}
