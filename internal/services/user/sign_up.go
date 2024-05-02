package user

import (
	"context"
	usererrors "github.com/DenisKhanov/PrivateKeeper/internal/errors"
	"github.com/DenisKhanov/PrivateKeeper/internal/models"
	"github.com/DenisKhanov/PrivateKeeper/pkg/auth"
	"github.com/DenisKhanov/PrivateKeeper/pkg/validate"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// SignUp handles the user registration process. It accepts a context (`ctx`) for managing request duration and a `User` model (`user`) containing the user's name, email, login, and password.
//
// The method first validates the provided registration data (login, password, email) using a validation function. If validation fails, it returns an error.
//
// It then checks if the login and email already exist in the database using `CheckExists` method. If either the login or email already exists, it returns a specific error (`ErrLoginAlreadyTaken` or `ErrEmailAlreadyTaken`).
//
// If the login and email are unique, the method hashes the user's password and generates a JWT token using the provided `auth` functions. If any errors occur during hashing or token generation, it logs the error and returns an appropriate error.
//
// If all steps are successful, the method saves the new user to the database using the `AddUser` method and returns the generated token.
//
// Returns:
// - The generated token string (`token`) if registration is successful.
// - An error if any step fails, including validation, checking for existing users, password hashing, token generation, or saving the user.
func (u *ServiceUser) SignUp(ctx context.Context, user models.User) (token string, err error) {
	if err = validate.CheckRegistrationData(user.Login, user.Password, user.Email); err != nil {
		logrus.Error(err)
		return "", err
	}
	//TODO уйти от дублирования кода
	exists, err := u.repository.CheckExists(ctx, user.Login)
	logrus.Info("CheckExists: ", exists)
	if err != nil {
		logrus.WithError(err).Error("Failed to check user existence")
		return "", err
	}

	if exists {
		logrus.Error(usererrors.ErrLoginAlreadyTaken)
		return "", usererrors.ErrLoginAlreadyTaken
	}

	exists, err = u.repository.CheckExists(ctx, user.Email)
	logrus.Info("CheckExists: ", exists)
	if err != nil {
		logrus.WithError(err).Error("Failed to check user existence")
		return "", err
	}

	if exists {
		logrus.Error(usererrors.ErrEmailAlreadyTaken)
		return "", usererrors.ErrEmailAlreadyTaken
	}

	hashedPassword, err := auth.CreateHashPassword(user.Password)
	if err != nil {
		logrus.WithError(err).Error("Failed to hash password")
		return "", err
	}
	//TODO нужно ли выходить с ошибкой или продолжить без токена
	userID := uuid.New()
	token, err = auth.BuildJWTString(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to build token")
		return "", usererrors.ErrSaveNewUser
	}
	if err = u.repository.AddUser(ctx, userID, user.Name, user.Email, user.Login, hashedPassword); err != nil {
		logrus.WithError(err).Error("Failed to save new user in database")
		return "", usererrors.ErrSaveNewUser
	}
	return token, nil
}

// withTransaction creates and manages a database transaction.
// If txFunc completes successfully, the transaction is committed.
// If txFunc returns an error or a panic occurs, the transaction is rolled back.
//
// Parameters:
// - ctx: The context for the transaction.
// - txFunc: A function that takes a transaction (pgx.Tx) as an argument.
//
// Returns:
// - An error if any issues occurred during the transaction or within txFunc.
func (u *ServiceUser) withTransaction(ctx context.Context, txFunc func(pgx.Tx) error) error {
	tx, err := u.dbPool.Begin(ctx)
	if err != nil {
		logrus.Error("Error starting transaction: ", err)
		return err
	}
	// transaction management using closure
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = txFunc(tx)
	return err
}
