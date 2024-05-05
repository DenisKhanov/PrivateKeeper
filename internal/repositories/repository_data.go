package repositories

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

//TODO изменить нейминг перегруппировать интерфесы интерфейсов

type S3Repository interface {
	AddBinaryData(ctx context.Context, data models.BinaryData) error
	GetBinaryData(ctx context.Context, objectName string) ([]byte, error)
	DelData(ctx context.Context, objectName string) error
}
type DataRepository interface {
	RepoLoginPasswordData
	RepoCardData
	RepoTextData
	RepoBinaryData
	RepoAllUserDataList
	RepoDataDeleter
}

type RepoLoginPasswordData interface {
	AddLoginPasswordData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.LoginData) error
	GetLoginPasswordData(ctx context.Context, userID uuid.UUID, metadataID int) (models.LoginData, error)
}

type RepoCardData interface {
	AddCardData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.CardData) error
	GetCardData(ctx context.Context, userID uuid.UUID, metadataID int) (models.CardData, error)
}

type RepoTextData interface {
	AddTextData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.TextData) error
	GetTextData(ctx context.Context, userID uuid.UUID, metadataID int) (models.TextData, error)
}

type RepoBinaryData interface {
	AddBinaryData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.BinaryData) error
	GetBinaryData(ctx context.Context, userID uuid.UUID, metadataID int) (models.BinaryData, error)
	GetS3ObjectName(ctx context.Context, userID uuid.UUID, metadataID int) (string, error)
}

type RepoAllUserDataList interface {
	GetAllUserDataList(ctx context.Context, userID uuid.UUID) ([]models.Metadata, error)
}

type RepoDataDeleter interface {
	DelData(ctx context.Context, userID uuid.UUID, metadataID int) error
}
