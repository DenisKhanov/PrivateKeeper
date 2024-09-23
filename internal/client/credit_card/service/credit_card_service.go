package service

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/client/state"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
)

type CreditCardService interface {
	SaveCreditCard(token string, card models.CardData) error
	LoadCreditCard(token string, metadataID uint64) (models.CardData, error)
}

type CreditCardProvider struct {
	creditCardService CreditCardService
	state             *state.ClientState
}

func NewUserService(u CreditCardService, state *state.ClientState) *CreditCardProvider {
	return &CreditCardProvider{
		creditCardService: u,
		state:             state,
	}
}
