package user

import (
	"context"
	"github.com/sirupsen/logrus"
)

// CheckExists checks whether a specified login or email exists in the database.
// It accepts a context (`ctx`) for managing request duration and a string (`data`) representing the login or email to check.
// The method queries the `users` table for a match in the `login` column using the provided data.
// It returns:
// - `true` if the login or email exists.
// - `false` otherwise.
func (u *RepositoryUser) CheckExists(ctx context.Context, data string) (bool, error) {
	var exists bool
	// Query to check for the existence of a login
	const selectQuery = `SELECT EXISTS (SELECT 1 FROM users WHERE login = $1 OR email = $1)`
	err := u.dbPool.QueryRow(ctx, selectQuery, data).Scan(&exists)
	if err != nil {
		logrus.WithError(err).Error("Failed to check login existence")
		return false, err
	}
	return exists, nil
}
