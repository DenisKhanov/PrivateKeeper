package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (u *RepositoryUser) GetUUID(ctx context.Context, login string) (userID uuid.UUID, err error) {
	const selectQuery = `SELECT uuid FROM users WHERE login = $1`
	err = u.dbPool.QueryRow(ctx, selectQuery, login).Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logrus.WithError(err).Error("Login not found")
			return uuid.Nil, err
		}
		logrus.WithError(err).Error("Error querying for uuid")
		return uuid.Nil, err
	}
	return userID, nil
}
