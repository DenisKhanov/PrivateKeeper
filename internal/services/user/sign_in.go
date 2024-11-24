package user

import (
	"context"
	usererrors "github.com/DenisKhanov/PrivateKeeper/internal/errors"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/DenisKhanov/PrivateKeeper/pkg/auth"
)

// SignIn handles the user sign-in process and returns a JWT token upon successful authentication.
// It accepts a context (`ctx`) for managing request duration and a `User` model (`user`) containing the user's login and password.
//
// The method retrieves the hashed password for the given user's login using the `GetPassword` method from the repository.
// If there is an error (e.g., the login does not exist), it returns an error indicating unauthorized access (`ErrUnauthorizedUser`).
//
// Next, it compares the saved hashed password with the password provided by the user using the `ComparePasswords` function from the `auth` package.
// If the comparison fails, it returns an error indicating unauthorized access (`ErrUnauthorizedUser`).
//
// If the passwords match, the method retrieves the user's UUID using the `GetUUID` method from the repository.
// If there is an error, it returns an error indicating issues with accessing the database (`ErrAccessingDB`).
//
// Finally, it generates a JWT token using the user's UUID with the `BuildJWTString` function from the `auth` package.
// If an error occurs during token generation, it returns an error indicating an issue with signing in (`ErrSignInUser`).
//
// Returns:
// - The generated JWT token (`token`) if sign-in is successful.
// - An error if any step fails, including retrieval of hashed password, password comparison, retrieval of user UUID, or token generation.
func (u *ServiceUser) SignIn(ctx context.Context, user models.User) (token string, err error) {
	savedHashedPassword, err := u.repository.GetPassword(ctx, user.Login)
	if err != nil {
		return "", usererrors.ErrUnauthorizedUser
	}

	if err = auth.ComparePasswords(savedHashedPassword, user.Password); err != nil {
		return "", usererrors.ErrUnauthorizedUser
	}

	savedUserID, err := u.repository.GetUUID(ctx, user.Login)
	if err != nil {
		return "", usererrors.ErrAccessingDB
	}

	token, err = auth.BuildJWTString(savedUserID)
	if err != nil {
		return "", usererrors.ErrSignInUser
	}

	return token, nil
}
