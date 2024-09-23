package pbclient

import (
	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

type CreditCardPBClient struct {
	creditCardService pb.KeeperDataV1Client
}

func NewCreditCardPBClient(u pb.KeeperDataV1Client) *CreditCardPBClient {
	return &CreditCardPBClient{
		creditCardService: u,
	}
}
