package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/domain"
	"github.com/google/uuid"
)

func (d *ServiceData) DelData(ctx context.Context, userID uuid.UUID, metadataID int, dataType string) error {
	if dataType == domain.BinaryData {
		objectName, err := d.repository.GetS3ObjectName(ctx, userID, metadataID)
		if err != nil {
			return err
		}
		if err = d.s3Repository.DelData(ctx, objectName); err != nil {
			return err
		}
	}
	if err := d.repository.DelData(ctx, userID, metadataID); err != nil {
		return err
	}
	return nil
}
