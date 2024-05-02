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
}

type ServiceLoginPassword interface {
	AddLoginPasswordData(ctx context.Context, userID uuid.UUID, data models.LoginData) error
}

type ServiceCardData interface {
	AddCardData(ctx context.Context, userID uuid.UUID, data models.CardData) error
}

type ServiceTextData interface {
	AddTextData(ctx context.Context, userID uuid.UUID, data models.TextData) error
}

type ServiceBinaryData interface {
	AddBinaryData(ctx context.Context, userID uuid.UUID, data models.BinaryData) error
}
