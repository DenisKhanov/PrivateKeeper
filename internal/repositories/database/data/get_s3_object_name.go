package data

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (d *RepositoryData) GetS3ObjectName(ctx context.Context, userID uuid.UUID, metadataID int) (string, error) {
	var objectName string
	const sqlQuery = `
		SELECT
    		binary_data.s3_object_name
		FROM
    		data_units
		JOIN
    		binary_data ON data_units.binary_data_id = binary_data.id
		WHERE
    		data_units.user_id = $1
    		AND data_units.metadata_id = $2; 
	`
	if err := d.dbPool.QueryRow(ctx, sqlQuery, userID, metadataID).Scan(&objectName); err != nil {
		logrus.WithError(err).Error("failed to get s3 object name")
		return "", err
	}
	logrus.Infof("Success got s3 object name: %s", objectName)
	return objectName, nil
}
