package data

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	proto "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (d *GRPCData) DelData(ctx context.Context, in *proto.DelDataRequest) (*proto.DelDataResponse, error) {
	userID, ok := ctx.Value(models.UserIDKey).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Internal, `could not find user ID in context`)
	}
	metadataID := int(in.MetadataId)
	dataType := in.DataType.String()
	if err := d.service.DelData(ctx, userID, metadataID, dataType); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return nil, status.Error(codes.OK, "")
}
