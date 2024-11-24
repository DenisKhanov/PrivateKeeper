package pbclient

import (
	"context"
	"google.golang.org/grpc/metadata"

	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

func (u *CreditCardPBClient) SaveCreditCard(token string, card models.CardData) error {
	// Создание DataUnit с типом BANK_CARD и заполнение данными банковской карты
	dataUnit := &pb.DataUnit{
		DataType: pb.DataType_BANK_CARD,
		Data: &pb.DataUnit_BankCard{
			BankCard: &pb.BankCard{
				HolderName:     card.HolderName,
				Number:         card.Number,
				ExpirationDate: card.ExpDate,
				Cvv:            card.CVV,
			},
		},
		MetaInfo: &pb.MetaInfo{
			Data: &pb.MetaInfo_Bank{
				Bank: card.Info,
			},
		},
	}

	// Создание запроса AddDataRequest с заполненным DataUnit
	req := &pb.AddDataRequest{
		DataUnits: []*pb.DataUnit{dataUnit},
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := u.creditCardService.AddData(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
