// Package errors defines common models and errors for the application.
package errors

import "errors"

var (
	ErrLoginAlreadyTaken = errors.New("login already taken")
	ErrEmailAlreadyTaken = errors.New("email already taken")
	ErrSaveNewUser       = errors.New("it is not possible to save the user to the database, try again")
	ErrSignInUser        = errors.New("it is not possible to sign in, try again")
	ErrUnauthorizedUser  = errors.New("login or password uncorrected")
	ErrAccessingDB       = errors.New("errors accessing the database")
)
