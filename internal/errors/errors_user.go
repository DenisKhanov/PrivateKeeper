// Package errors defines common models and errors for the application.
package errors

import "errors"

var (
	ErrLoginAlreadyTaken      = errors.New("login already taken")
	ErrEmailAlreadyTaken      = errors.New("email already taken")
	ErrSaveNewUser            = errors.New("it is not possible to save the user to the database, try again")
	ErrSignInUser             = errors.New("it is not possible to sign in, try again")
	ErrUnauthorizedUser       = errors.New("login or password uncorrected")
	ErrTokenIsNotValid        = errors.New("token is not valid")
	ErrOrderNumber            = errors.New("invalid order number format")
	ErrUserOrderExists        = errors.New("the order number has already been uploaded by this user")
	ErrAnotherUserOrderExists = errors.New("the order number has already been uploaded by another user")
	ErrAccessingDB            = errors.New("errors accessing the database")
	ErrUserHasNoOrders        = errors.New("this user does not have any orders")
	ErrUserHasNoWithdrawals   = errors.New("this user does not have any withdrawals")
	ErrNotEnoughFunds         = errors.New("there are not enough funds in the bonus account to be debited")
)
