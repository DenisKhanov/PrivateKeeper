package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/domain"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	proto "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (d *GRPCData) AddData(ctx context.Context, in *proto.AddDataRequest) (*proto.AddDataResponse, error) {
	//TODO исправить падение сервиса при получении не верного типа в запросе от пользователя
	userID, ok := ctx.Value(domain.UserIDKey).(uuid.UUID)
	if !ok {
		logrus.Info("Could not extract UserID from ctx")
		return nil, status.Error(codes.Internal, "could not find user ID in context")
	}
	var (
		loginData  models.LoginData
		cardData   models.CardData
		textData   models.TextData
		binaryData models.BinaryData
	)
	logrus.Info(proto.DataType_LOGIN_PASSWORD)
	for _, dataUnit := range in.DataUnits {
		switch dataUnit.DataType {
		case proto.DataType_LOGIN_PASSWORD:
			loginData.DataType = dataUnit.DataType.String()
			loginData.Login = dataUnit.GetLoginPassword().Login
			loginData.Password = dataUnit.GetLoginPassword().Password
			loginData.Info = dataUnit.MetaInfo.GetWebsite()
			if loginData.Login == "" && loginData.Password == "" {
				return nil, status.Error(codes.InvalidArgument, "login and password can't be empty")
			}
			if err := d.service.AddLoginPasswordData(ctx, userID, loginData); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		case proto.DataType_BANK_CARD:
			cardData.DataType = dataUnit.DataType.String()
			cardData.CVV = dataUnit.GetBankCard().Cvv
			cardData.Number = dataUnit.GetBankCard().Number
			cardData.ExpDate = dataUnit.GetBankCard().ExpirationDate
			cardData.HolderName = dataUnit.GetBankCard().HolderName
			cardData.Info = dataUnit.MetaInfo.GetBank()
			if cardData.CVV == "" && cardData.Number == "" && cardData.ExpDate == "" && cardData.HolderName == "" {
				return nil, status.Error(codes.InvalidArgument, "card's data can't be empty")
			}
			if err := d.service.AddCardData(ctx, userID, cardData); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		case proto.DataType_TEXT_DATA:
			textData.DataType = dataUnit.DataType.String()
			textData.Content = dataUnit.GetTextData().Content
			textData.Info = dataUnit.MetaInfo.GetTextDataDescription()
			if textData.Content == "" {
				return nil, status.Error(codes.InvalidArgument, "text data can't be empty")
			}
			if err := d.service.AddTextData(ctx, userID, textData); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		case proto.DataType_BINARY_DATA:
			binaryData.DataType = dataUnit.DataType.String()
			binaryData.ObjectName = dataUnit.GetBinaryData().ObjectName
			binaryData.Content = dataUnit.GetBinaryData().Content
			binaryData.Info = dataUnit.MetaInfo.GetBinaryDataDescription()
			if binaryData.ObjectName == "" && binaryData.Content == nil {
				return nil, status.Error(codes.InvalidArgument, "binary data and object name can't be empty")
			}
			if err := d.service.AddBinaryData(ctx, userID, binaryData); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}

		default:
			return nil, status.Errorf(codes.InvalidArgument, "Unsupported data type: %v", dataUnit.DataType)
		}
	}

	return nil, status.Errorf(codes.OK, "Data successfully added")
}
