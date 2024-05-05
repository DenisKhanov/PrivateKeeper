package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// AddUser saves a new user with a pre-generated UUID, name, login, and hashed password.
func (u *RepositoryUser) AddUser(ctx context.Context, userID uuid.UUID,
	name, email, login string, hashedPassword []byte) error {

	const sqlQuery = `INSERT INTO users (uuid,name,email,login, password_hash) 
					  VALUES ($1,$2,$3,$4,$5) 
					  ON CONFLICT (login) DO NOTHING
`
	_, err := u.dbPool.Exec(ctx, sqlQuery, userID, name, email, login, hashedPassword)
	if err != nil {
		logrus.WithError(err).Error("Failed to save new user in database")
		return err
	}
	return nil
}
