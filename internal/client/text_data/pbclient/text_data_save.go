package pbclient

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"

	"google.golang.org/grpc/metadata"

	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

func (u *TextDataPBClient) SaveTextData(token string, text models.TextData) error {
	dataUnit := &pb.DataUnit{
		DataType: pb.DataType_TEXT_DATA,
		Data: &pb.DataUnit_TextData{
			TextData: &pb.TextData{
				Content: text.Content,
			},
		},
		MetaInfo: &pb.MetaInfo{
			Data: &pb.MetaInfo_TextDataDescription{
				TextDataDescription: text.Info,
			},
		},
	}

	// Создание запроса AddDataRequest с заполненным DataUnit
	req := &pb.AddDataRequest{
		DataUnits: []*pb.DataUnit{dataUnit},
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := u.textDataService.AddData(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
