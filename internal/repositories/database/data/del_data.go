package data

import (
	"context"
	errors3 "github.com/DenisKhanov/PrivateKeeper/internal/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (d *RepositoryData) DelData(ctx context.Context, userID uuid.UUID, metadataID int) error {
	const sqlQuery = `
        DELETE FROM data_units
        WHERE user_id = $1
            AND metadata_id = $2;
    `
	_, err := d.dbPool.Exec(ctx, sqlQuery, userID, metadataID)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete data.")
		return errors3.ErrDelData
	}
	logrus.Info("Successfully deleted data.")
	return nil
}
