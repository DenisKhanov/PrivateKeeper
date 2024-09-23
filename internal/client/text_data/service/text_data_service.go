package service

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/client/state"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
)

type TextDataService interface {
	SaveTextData(token string, text models.TextData) error
	LoadTextData(token string, metadataID uint64) (models.TextData, error)
}

type TextDataProvider struct {
	textDataService TextDataService
	state           *state.ClientState
}

func NewTextDataService(u TextDataService, state *state.ClientState) *TextDataProvider {
	return &TextDataProvider{
		textDataService: u,
		state:           state,
	}
}
