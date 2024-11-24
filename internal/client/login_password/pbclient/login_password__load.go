package pbclient

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"

	"google.golang.org/grpc/metadata"

	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

func (u *LoginPasswordPBClient) LoadLoginPassword(token string, metadataID uint64) (models.LoginData, error) {
	req := &pb.GetDataRequest{
		DataType:   pb.DataType_LOGIN_PASSWORD,
		MetadataId: metadataID,
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.loginPasswordService.GetData(ctx, req)
	if err != nil {
		return models.LoginData{}, err
	}
	respLoginData := resp.DataUnit
	loginData := models.LoginData{
		Login:    respLoginData.GetLoginPassword().Login,
		Password: respLoginData.GetLoginPassword().Password,
	}

	return loginData, nil
}
