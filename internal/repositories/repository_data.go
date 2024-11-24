package repositories

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/minio/minio-go/v7"
)

//TODO изменить нейминг, перегруппировать интерфесы

type S3Repository interface {
	AddBinaryData(ctx context.Context, data models.EncryptedBinaryData, encryptedContent []byte) (minio.UploadInfo, error)
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
	RepoKeyManager
}

type RepoLoginPasswordData interface {
	AddLoginPasswordData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.KeepData) error
	GetLoginPasswordData(ctx context.Context, userID uuid.UUID, metadataID int) (models.KeepData, error)
}

type RepoCardData interface {
	AddCardData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.KeepData) error
	GetCardData(ctx context.Context, userID uuid.UUID, metadataID int) (models.KeepData, error)
}

type RepoTextData interface {
	AddTextData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.KeepData) error
	GetTextData(ctx context.Context, userID uuid.UUID, metadataID int) (models.KeepData, error)
}

type RepoBinaryData interface {
	AddBinaryData(ctx context.Context, tx pgx.Tx, userID uuid.UUID, data models.EncryptedBinaryData) error
	GetBinaryData(ctx context.Context, userID uuid.UUID, metadataID int) (models.EncryptedBinaryData, error)
	GetS3ObjectName(ctx context.Context, userID uuid.UUID, metadataID int) (string, error)
}

type RepoAllUserDataList interface {
	GetAllUserDataList(ctx context.Context, userID uuid.UUID) ([]models.Metadata, error)
}

type RepoDataDeleter interface {
	DelData(ctx context.Context, userID uuid.UUID, metadataID int) error
}

type RepoKeyManager interface {
	GetEncryptedKey(ctx context.Context, userID uuid.UUID) ([]byte, error)
}
