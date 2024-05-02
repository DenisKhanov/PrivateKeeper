package repositories

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type S3Repository interface {
	AddBinaryData(ctx context.Context, objectName string, data []byte) (string, error)
	GetBinaryData(ctx context.Context, objectName string) ([]byte, error)
}
type DataRepository interface {
	RepoLoginPassword
	RepoCardData
	RepoTextData
	RepoBinaryData
}

type RepoLoginPassword interface {
	AddLoginPasswordData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.LoginData) error
}

type RepoCardData interface {
	AddCardData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.CardData) error
}

type RepoTextData interface {
	AddTextData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.TextData) error
}

type RepoBinaryData interface {
	AddBinaryData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, dataInfo, s3URL string) error
}
