package pbclient

import (
	"context"

	"google.golang.org/grpc/metadata"

	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

func (u *DataPBClient) LoadAllDataInfo(token string) ([]*pb.DataInfo, error) {
	req := &pb.AllUserDataListRequest{}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.dataService.AllUserDataList(ctx, req)
	if err != nil {
		return nil, err
	}
	var data pb.DataInfo
	allMetadata := make([]*pb.DataInfo, len(resp.AllDataList))
	for _, i := range resp.AllDataList {
		data = pb.DataInfo{
			DataId:      i.DataId,
			DataType:    i.DataType,
			Description: i.Description,
		}
		allMetadata = append(allMetadata, &data)
	}
	return allMetadata, nil
}
