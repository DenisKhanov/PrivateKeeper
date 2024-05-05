package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
)

func (d *ServiceData) AllUserDataList(ctx context.Context, userID uuid.UUID) ([]models.Metadata, error) {
	metadataList, err := d.repository.GetAllUserDataList(ctx, userID)
	if err != nil {
		return nil, err
	}
	return metadataList, nil
}
