package service

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/client/state"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
)

type LoginPasswordService interface {
	SaveLoginPassword(token string, cred models.LoginData) error
	LoadLoginPassword(token string, metadataID uint64) (models.LoginData, error)
}

type LoginPasswordProvider struct {
	loginPasswordService LoginPasswordService
	state                *state.ClientState
}

func NewCredentialsService(u LoginPasswordService, state *state.ClientState) *LoginPasswordProvider {
	return &LoginPasswordProvider{
		loginPasswordService: u,
		state:                state,
	}
}
