package user

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/services"
	serviceuser "github.com/DenisKhanov/PrivateKeeper/internal/services/user"
	protouser "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/user"
)

// checking interface compliance at the compiler level
var _ services.UserService = (*serviceuser.ServiceUser)(nil)

// GRPCUser ...
type GRPCUser struct {
	protouser.UnimplementedKeeperUserV1Server //for compatibility with new versions
	service                                   services.UserService
}

// NewGRPCUser ...
func NewGRPCUser(service services.UserService) *GRPCUser {
	return &GRPCUser{
		service: service,
	}
}
