package services

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/google/uuid"
)

type DataService interface {
	ServiceLoginPassword
	ServiceCardData
	ServiceTextData
	ServiceBinaryData
	ServiceAllUserDataList
}

type ServiceLoginPassword interface {
	AddLoginPasswordData(ctx context.Context, userID uuid.UUID, data models.LoginData) error
	GetDecodedLoginPasswordData(ctx context.Context, userID uuid.UUID, metadataID int) (models.LoginData, error)
}

type ServiceCardData interface {
	AddCardData(ctx context.Context, userID uuid.UUID, data models.CardData) error
	GetDecodedBankCardData(ctx context.Context, userID uuid.UUID, metadataID int) (models.CardData, error)
}

type ServiceTextData interface {
	AddTextData(ctx context.Context, userID uuid.UUID, data models.TextData) error
	GetDecodedTextData(ctx context.Context, userID uuid.UUID, metadataID int) (models.TextData, error)
}

type ServiceBinaryData interface {
	AddBinaryData(ctx context.Context, userID uuid.UUID, data models.BinaryData) error
	GetDecodedBinaryData(ctx context.Context, userID uuid.UUID, metadataID int) (models.BinaryData, error)
}

type ServiceAllUserDataList interface {
	AllUserDataList(ctx context.Context, userID uuid.UUID) ([]models.Metadata, error)
}
