package pbclient

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"

	"google.golang.org/grpc/metadata"

	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

func (u *TextDataPBClient) LoadTextData(token string, metadataID uint64) (models.TextData, error) {
	req := &pb.GetDataRequest{
		DataType:   pb.DataType_TEXT_DATA,
		MetadataId: metadataID,
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.textDataService.GetData(ctx, req)
	if err != nil {
		return models.TextData{}, err
	}
	respTextData := resp.DataUnit
	textData := models.TextData{
		Content: respTextData.GetTextData().Content,
	}

	return textData, nil
}
