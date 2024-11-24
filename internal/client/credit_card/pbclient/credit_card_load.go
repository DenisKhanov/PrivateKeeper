package pbclient

import (
	"context"

	"google.golang.org/grpc/metadata"

	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

func (u *CreditCardPBClient) LoadCreditCard(token string, metadataID uint64) (models.CardData, error) {
	req := &pb.GetDataRequest{
		DataType:   pb.DataType_BANK_CARD,
		MetadataId: metadataID,
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.creditCardService.GetData(ctx, req)
	if err != nil {
		return models.CardData{}, err
	}
	respCardData := resp.DataUnit
	card := models.CardData{
		CVV:        respCardData.GetBankCard().Cvv,
		Number:     respCardData.GetBankCard().Number,
		HolderName: respCardData.GetBankCard().HolderName,
		ExpDate:    respCardData.GetBankCard().ExpirationDate,
	}

	return card, nil
}
