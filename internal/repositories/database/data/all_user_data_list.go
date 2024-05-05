package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (d *RepositoryData) GetAllUserDataList(ctx context.Context, userID uuid.UUID) ([]models.Metadata, error) {
	const sqlQuery = `SELECT
    				  		metadata.id AS metadata_id,
    						metadata.data_type,
    						COALESCE(metadata.website, metadata.bank, metadata.text_data_description, metadata.binary_data_description) AS metadata_description
					FROM
    						data_units
					JOIN
    						metadata ON data_units.metadata_id = metadata.id
					WHERE
    						data_units.user_id = $1
 					AND (
       						metadata.website IS NOT NULL AND metadata.website <> ''
        					OR metadata.bank IS NOT NULL AND metadata.bank <> ''
        					OR metadata.text_data_description IS NOT NULL AND metadata.text_data_description <> ''
        					OR metadata.binary_data_description IS NOT NULL AND metadata.binary_data_description <> ''
    				);`
	rows, err := d.dbPool.Query(ctx, sqlQuery, userID)
	if err != nil {
		logrus.WithError(err).Error("Metadata could not be retrieved")
		return nil, err
	}
	defer rows.Close()

	var result []models.Metadata
	for rows.Next() {
		var metadata models.Metadata
		if err = rows.Scan(&metadata.ID, &metadata.DataType, &metadata.Description); err != nil {
			logrus.WithError(err).Error("Metadata could not be retrieved")
			return nil, err
		}
		result = append(result, metadata)
	}
	if err = rows.Err(); err != nil {
		logrus.WithError(err).Error("Error occurred during reading rows")
		return nil, err
	}
	logrus.Infof("Successfully retrieved all user data list: %v", result)
	return result, nil
}
