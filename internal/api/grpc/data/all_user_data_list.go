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

func (d *GRPCData) AllUserDataList(ctx context.Context,
	_ *proto.AllUserDataListRequest) (*proto.AllUserDataListResponse, error) {
	userID, ok := ctx.Value(models.UserIDKey).(uuid.UUID)
	if !ok {
		logrus.Info("Could not extract UserID from ctx")
		return nil, status.Error(codes.Internal, "could not find user ID in context")
	}
	metadataList, err := d.service.AllUserDataList(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, "could not list all user data")
	}
	response := &proto.AllUserDataListResponse{
		AllDataList: make([]*proto.DataInfo, len(metadataList)),
	}
	for i, metadata := range metadataList {
		response.AllDataList[i] = &proto.DataInfo{
			DataId:      uint64(metadata.ID),
			DataType:    metadata.DataType,
			Description: metadata.Description,
		}
	}
	return response, nil
}
