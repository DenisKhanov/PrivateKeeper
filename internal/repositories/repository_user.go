package repositories

import (
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	CheckExists(ctx context.Context, data string) (bool, error)
	AddUser(ctx context.Context, userID uuid.UUID,
		name, email, login string, hashedPassword []byte) error
	GetPassword(ctx context.Context, login string) (hashedPassword []byte, err error)
	GetUUID(ctx context.Context, login string) (userId uuid.UUID, err error)
}
