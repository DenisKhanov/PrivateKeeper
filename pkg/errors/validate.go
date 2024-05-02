package errors

import "errors"

var (
	ErrLoginShort = errors.New("the login is too short")
	ErrLoginLong  = errors.New("the login is too long")
	ErrLoginChar  = errors.New("the login contains invalid characters")
)

var (
	ErrPasswordMinLen  = errors.New("the password is too short")
	ErrPasswordMaxLen  = errors.New("the password is too long")
	ErrPasswordUpper   = errors.New("the password must contain at least one capital letter")
	ErrPasswordLower   = errors.New("the password must contain at least one lowercase letter")
	ErrPasswordDigit   = errors.New("the password must contain at least one digit")
	ErrPasswordSpecial = errors.New("the password must contain at least one special character")
)

var (
	ErrEmail = errors.New("invalid email")
)
