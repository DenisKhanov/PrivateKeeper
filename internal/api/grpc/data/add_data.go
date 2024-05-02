package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	proto "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (d *GRPCData) AddData(ctx context.Context, in *proto.AddDataRequest) (*proto.AddDataResponse, error) {
	userID, ok := ctx.Value(models.UserIDKey).(uuid.UUID)
	if !ok {
		logrus.Errorf("context value is not userID: %v", userID)
	}
	logrus.Infof("AddData UUID: %v", userID)
	var (
		loginData  models.LoginData
		cardData   models.CardData
		textData   models.TextData
		binaryData models.BinaryData
	)
	for _, dataUnit := range in.DataUnits {
		switch dataUnit.Type {
		case proto.DataType_LOGIN_PASSWORD:
			loginData.Login = dataUnit.GetLoginPassword().Login
			loginData.Password = dataUnit.GetLoginPassword().Password
			loginData.Info = dataUnit.MetaInfo.GetWebsite()
			if err := d.service.AddLoginPasswordData(ctx, userID, loginData); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		case proto.DataType_BANK_CARD:
			cardData.CVV = dataUnit.GetBankCard().Cvv
			cardData.Number = dataUnit.GetBankCard().Number
			cardData.ExpDate = dataUnit.GetBankCard().ExpirationDate
			cardData.HolderName = dataUnit.GetBankCard().HolderName
			cardData.Info = dataUnit.MetaInfo.GetBank()
			if err := d.service.AddCardData(ctx, userID, cardData); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		case proto.DataType_TEXT_DATA:
			textData.Content = dataUnit.GetTextData().Content
			textData.Info = dataUnit.MetaInfo.GetTextDataDescription()
			if err := d.service.AddTextData(ctx, userID, textData); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		case proto.DataType_BINARY_DATA:

			binaryData.ObjectName = dataUnit.GetBinaryData().GetName()
			binaryData.Content = dataUnit.GetBinaryData().GetContent()
			binaryData.Info = dataUnit.MetaInfo.GetBinaryDataDescription()
			if err := d.service.AddBinaryData(ctx, userID, binaryData); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		default:
			return nil, status.Errorf(codes.InvalidArgument, "Unsupported data type: %v", dataUnit.Type)
		}
	}

	return nil, status.Errorf(codes.OK, "Data successfully added")
}
