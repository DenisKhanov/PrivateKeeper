package services

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
)

//go:generate mockgen -source=grpc.go -destination=mocks/grpc_mock.go -package=mocks
type UserService interface {
	SignUp(ctx context.Context, user models.User) (token string, err error)
	SignIn(ctx context.Context, user models.User) (token string, err error)
}
