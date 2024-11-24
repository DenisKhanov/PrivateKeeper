package user

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// GetPassword retrieves the hashed password for a user with the given login from the database.
// It accepts a context (`ctx`) for managing request duration and a string (`login`) representing the user's login.
//
// The method executes a query to retrieve the `password_hash` from the `users` table for the specified login.
// If the query is successful, the hashed password is returned.
// If the specified login does not exist in the database, the method logs an error and returns an error indicating that the login was not found (`pgx.ErrNoRows`).
// If there is any other error during the query, it logs the error and returns the error.
//
// Returns:
// - The hashed password as a byte slice (`hashedPassword`) if the query is successful.
// - An error if the specified login does not exist or if another error occurs during the query.
func (u *RepositoryUser) GetPassword(ctx context.Context, login string) (hashedPassword []byte, err error) {
	const selectQuery = `SELECT password_hash FROM users WHERE login = $1`
	if err = u.dbPool.QueryRow(ctx, selectQuery, login).Scan(&hashedPassword); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithError(err).Error("Login not found")
			return nil, err
		}
		logrus.WithError(err).Error("Error querying for savedHashedPassword")
		return nil, err
	}
	return hashedPassword, nil
}
