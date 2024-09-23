package data

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (d *RepositoryData) GetEncryptedKey(ctx context.Context, userID uuid.UUID) ([]byte, error) {
	var encryptedKey []byte
	const sqlQuery = `
		SELECT
    		encrypted_key
		FROM
    		users
		WHERE
    		uuid = $1; 
	`
	if err := d.dbPool.QueryRow(ctx, sqlQuery, userID).Scan(&encryptedKey); err != nil {
		logrus.WithError(err).Error("Error getting encrypted key")
		return nil, err
	}
	logrus.Info("Success got encryptedKey.")
	return encryptedKey, nil
}
