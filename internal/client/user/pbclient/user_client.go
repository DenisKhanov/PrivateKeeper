package pbclient

import (
	"context"

	pb1 "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/user"
)

type UserPBClient struct {
	userService pb1.KeeperUserV1Client
}

func NewUserPBClient(u pb1.KeeperUserV1Client) *UserPBClient {
	return &UserPBClient{
		userService: u,
	}
}

func (u *UserPBClient) LoginUser(login, password string) (string, error) {
	req := &pb1.SignInRequest{
		Login:    login,
		Password: password,
	}

	resp, err := u.userService.SignIn(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (u *UserPBClient) RegisterUser(name, email, login, password string) (string, error) {
	req := &pb1.SignUpRequest{
		Name:     name,
		Email:    email,
		Login:    login,
		Password: password,
	}

	resp, err := u.userService.SignUp(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}
