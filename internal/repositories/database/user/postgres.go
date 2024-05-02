package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// RepositoryUser ...
type RepositoryUser struct {
	dbPool *pgxpool.Pool //opened in main func dbPool pool connections
}

// TODO есть сомнения с правильностью создания таблиц в методе

func NewPostgresUser(dbPool *pgxpool.Pool) (*RepositoryUser, error) {
	storage := &RepositoryUser{
		dbPool: dbPool,
	}
	if err := storage.createTables(); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return storage, nil
}

// createTables ...
func (u *RepositoryUser) createTables() error {
	logrus.Infof("Creating table users")
	ctx := context.Background()
	sqlQuery := `
CREATE TABLE IF NOT EXISTS users (
    uuid UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash BYTEA NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS metadata (
    id SERIAL PRIMARY KEY,
    data_type VARCHAR(255) NOT NULL,
    website VARCHAR(255),
    bank VARCHAR(255),
    text_data_description VARCHAR(255),
    binary_data_description VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS logins_passwords (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS text_data (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS binary_data (
    id SERIAL PRIMARY KEY,
    s3_url VARCHAR(2083) NOT NULL
);
CREATE TABLE IF NOT EXISTS bank_cards (
    id SERIAL PRIMARY KEY,
    holder_name VARCHAR(255) NOT NULL,
    number VARCHAR(255) NOT NULL,
    expiration_date VARCHAR(255) NOT NULL,
    cvv VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS data_units (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(uuid) ON DELETE CASCADE,
    login_password_id INTEGER REFERENCES logins_passwords(id),
    text_data_id INTEGER REFERENCES text_data(id),
    binary_data_id INTEGER REFERENCES binary_data(id),
    bank_card_id INTEGER REFERENCES bank_cards(id),
    metadata_id INTEGER NOT NULL REFERENCES metadata(id)
);`
	_, err := u.dbPool.Exec(ctx, sqlQuery)
	if err != nil {
		logrus.WithError(err).Error("don't create tables:")
		return err
	}
	logrus.Info("Successfully created tables!")
	return nil
}
