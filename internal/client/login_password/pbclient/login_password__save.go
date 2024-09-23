package pbclient

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"

	"google.golang.org/grpc/metadata"

	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

func (u *LoginPasswordPBClient) SaveLoginPassword(token string, cred models.LoginData) error {
	dataUnit := &pb.DataUnit{
		DataType: pb.DataType_LOGIN_PASSWORD,
		Data: &pb.DataUnit_LoginPassword{
			LoginPassword: &pb.LoginPassword{
				Login:    cred.Login,
				Password: cred.Password,
			},
		},
		MetaInfo: &pb.MetaInfo{
			Data: &pb.MetaInfo_Website{
				Website: cred.Info,
			},
		},
	}
	// Создание запроса AddDataRequest с заполненным DataUnit
	req := &pb.AddDataRequest{
		DataUnits: []*pb.DataUnit{dataUnit},
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := u.loginPasswordService.AddData(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
