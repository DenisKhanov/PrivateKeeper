package data

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/services"
	servicedata "github.com/DenisKhanov/PrivateKeeper/internal/services/data"
	protodata "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

// checking interface compliance at the compiler level
var _ services.DataService = (*servicedata.ServiceData)(nil)

// GRPCData ...
type GRPCData struct {
	protodata.UnimplementedKeeperDataV1Server //for compatibility with new versions
	service                                   services.DataService
}

// NewGRPCData ...
func NewGRPCData(service services.DataService) *GRPCData {
	return &GRPCData{
		service: service,
	}
}
