package pbclient

import (
	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

type LoginPasswordPBClient struct {
	loginPasswordService pb.KeeperDataV1Client
}

func NewCredentialsPBClient(u pb.KeeperDataV1Client) *LoginPasswordPBClient {
	return &LoginPasswordPBClient{
		loginPasswordService: u,
	}
}
