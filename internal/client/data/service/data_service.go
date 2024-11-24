package service

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/client/state"
	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

type DataService interface {
	LoadAllDataInfo(token string) ([]*pb.DataInfo, error)
}

type DataProvider struct {
	dataService DataService
	state       *state.ClientState
}

func NewDataService(u DataService, state *state.ClientState) *DataProvider {
	return &DataProvider{
		dataService: u,
		state:       state,
	}
}
