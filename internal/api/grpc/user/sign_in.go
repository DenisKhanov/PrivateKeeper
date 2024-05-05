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

// SignIn handles a gRPC request for user authentication and returns a response.
// Accepts a context (`ctx`) and a `SignInRequest` (`in`) containing the user's login and password.
//
// Creates a `User` model from the request data and uses the `service.SignIn` method to authenticate the user.
// If authentication fails, returns a gRPC error with the appropriate status code and message.
//
// Sends the JWT token as metadata in the response header using `grpc.SendHeader`.
// If there is an error sending the token, returns a gRPC error with the appropriate status code and message.
//
// Returns a gRPC response with an `OK` status code and a success message.
func (u GRPCUser) SignIn(ctx context.Context, in *proto.SignInRequest) (*proto.SignInResponse, error) {
	var User = models.User{
		Login:    in.Login,
		Password: in.Password,
	}
	tokenString, err := u.service.SignIn(ctx, User)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error()) // Custom errors with explanations are displayed here

	}
	if err = grpc.SendHeader(ctx, metadata.New(map[string]string{
		string(models.TokenKey): tokenString})); err != nil {
		return nil, status.Errorf(codes.Unknown, `error send token in metadata: %v`, err)
	}
	return nil, status.Error(codes.OK, "Sign in successfully")
}
