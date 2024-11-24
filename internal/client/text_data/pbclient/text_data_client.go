package pbclient

import (
	pb "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/data"
)

type TextDataPBClient struct {
	textDataService pb.KeeperDataV1Client
}

func NewCreditCardPBClient(u pb.KeeperDataV1Client) *TextDataPBClient {
	return &TextDataPBClient{
		textDataService: u,
	}
}
