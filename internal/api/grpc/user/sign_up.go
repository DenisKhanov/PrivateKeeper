package user

import (
	"context"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	proto "github.com/DenisKhanov/PrivateKeeper/pkg/keeper_v1/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// SignUp registers a new user with the provided name, email, login, and password.
// It takes a context (`ctx`) for managing request duration and a `SignUpRequest` input (`in`) containing the user's name, email, login, and password.
//
// The method creates a `User` model using the input data and passes it to the `SignUp` service method.
// If successful, a token string is returned, which is then sent as metadata in the gRPC response header.
//
// Returns:
// - `nil` and an `OK` status if the user is successfully created.
// - An error and appropriate status codes (`InvalidArgument` or `Unknown`) if there are issues with the request or metadata.
func (u GRPCUser) SignUp(ctx context.Context, in *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	var newUser = models.User{
		Name:     in.Name,
		Email:    in.Email,
		Login:    in.Login,
		Password: in.Password,
	}
	tokenString, err := u.service.SignUp(ctx, newUser)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // Custom errors with explanations are displayed here

	}
	if err = grpc.SendHeader(ctx, metadata.New(map[string]string{
		string(models.TokenKey): tokenString})); err != nil {
		return nil, status.Errorf(codes.Unknown, `error send token in metadata: %v`, err)
	}
	return nil, status.Error(codes.OK, "User created successfully")
}
