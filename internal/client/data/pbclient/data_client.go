package pbclient

import (
	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

type DataPBClient struct {
	dataService pb.KeeperDataV1Client
}

func NewDataPBClient(u pb.KeeperDataV1Client) *DataPBClient {
	return &DataPBClient{
		dataService: u,
	}
}
