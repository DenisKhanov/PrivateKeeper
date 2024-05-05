package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (d *ServiceData) AddLoginPasswordData(ctx context.Context, userID uuid.UUID, data models.LoginData) error {
	if err := d.withTransaction(ctx, func(tx pgx.Tx) error {
		if err := d.repository.AddLoginPasswordData(ctx, tx, userID, data); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (d *ServiceData) AddCardData(ctx context.Context, userID uuid.UUID, data models.CardData) error {
	if err := d.withTransaction(ctx, func(tx pgx.Tx) error {
		if err := d.repository.AddCardData(ctx, tx, userID, data); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (d *ServiceData) AddTextData(ctx context.Context, userID uuid.UUID, data models.TextData) error {
	if err := d.withTransaction(ctx, func(tx pgx.Tx) error {
		if err := d.repository.AddTextData(ctx, tx, userID, data); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (d *ServiceData) AddBinaryData(ctx context.Context, userID uuid.UUID, data models.BinaryData) error {
	if err := d.s3Repository.AddBinaryData(ctx, data); err != nil {
		return err
	}
	newData := models.BinaryData{
		DataType:   data.DataType,
		ObjectName: data.ObjectName,
		Content:    nil,
		Info:       data.Info,
	}
	if err := d.withTransaction(ctx, func(tx pgx.Tx) error {
		if err := d.repository.AddBinaryData(ctx, tx, userID, newData); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
