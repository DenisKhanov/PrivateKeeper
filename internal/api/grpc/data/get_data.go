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

func (d *GRPCData) GetData(ctx context.Context, in *proto.GetDataRequest) (*proto.GetDataResponse, error) {
	userID, ok := ctx.Value(models.UserIDKey).(uuid.UUID)
	if !ok {
		logrus.Info("Could not extract UserID from ctx")
		return nil, status.Error(codes.Internal, "could not find user ID in context")
	}
	var response *proto.GetDataResponse

	metadataID := int(in.MetadataId)
	dataType := in.DataType

	switch dataType {
	case proto.DataType_LOGIN_PASSWORD:
		decodedLoginPasswordData, err := d.service.GetDecodedLoginPasswordData(ctx, userID, metadataID)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		response = &proto.GetDataResponse{
			DataUnit: &proto.DataUnit{
				DataType: dataType,
				Data: &proto.DataUnit_LoginPassword{
					LoginPassword: &proto.LoginPassword{
						Login:    decodedLoginPasswordData.Login,
						Password: decodedLoginPasswordData.Password,
					},
				},
				MetaInfo: &proto.MetaInfo{
					Data: &proto.MetaInfo_Website{
						Website: decodedLoginPasswordData.Info,
					},
				},
			},
		}
	case proto.DataType_BANK_CARD:
		bankCardData, err := d.service.GetDecodedBankCardData(ctx, userID, metadataID)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		response = &proto.GetDataResponse{
			DataUnit: &proto.DataUnit{
				DataType: dataType,
				Data: &proto.DataUnit_BankCard{
					BankCard: &proto.BankCard{
						HolderName:     bankCardData.HolderName,
						Number:         bankCardData.Number,
						ExpirationDate: bankCardData.ExpDate,
						Cvv:            bankCardData.CVV,
					},
				},
				MetaInfo: &proto.MetaInfo{
					Data: &proto.MetaInfo_Bank{
						Bank: bankCardData.Info,
					},
				},
			},
		}
	case proto.DataType_TEXT_DATA:
		textData, err := d.service.GetDecodedTextData(ctx, userID, metadataID)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		response = &proto.GetDataResponse{
			DataUnit: &proto.DataUnit{
				DataType: dataType,
				Data: &proto.DataUnit_TextData{
					TextData: &proto.TextData{
						Content: textData.Content,
					},
				},
				MetaInfo: &proto.MetaInfo{
					Data: &proto.MetaInfo_TextDataDescription{
						TextDataDescription: textData.Info,
					},
				},
			},
		}
	case proto.DataType_BINARY_DATA:
		binaryData, err := d.service.GetDecodedBinaryData(ctx, userID, metadataID)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		response = &proto.GetDataResponse{
			DataUnit: &proto.DataUnit{
				DataType: dataType,
				Data: &proto.DataUnit_BinaryData{
					BinaryData: &proto.BinaryData{
						ObjectName: binaryData.ObjectName,
						Content:    binaryData.Content,
					},
				},
				MetaInfo: &proto.MetaInfo{
					Data: &proto.MetaInfo_BinaryDataDescription{
						BinaryDataDescription: binaryData.Info,
					},
				},
			},
		}
	default:
		return nil, status.Error(codes.InvalidArgument, `data type not supported`)
	}

	return response, status.Error(codes.OK, "")
}
