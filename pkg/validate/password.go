package validate

import (
	errors2 "github.com/DenisKhanov/PrivateKeeper/pkg/errors"
	"unicode"
)

// CheckPassword method checks the password against the criteria (length must be at least minLen and no more than maxLen characters,
// contains at least one lowercase letter, at least one digit, and at least one special character)
func CheckPassword(password string, minLen, maxLen uint8) error {
	var (
		hasMinLen  = false
		hasMaxLen  = false
		hasUpper   = false
		hasLower   = false
		hasDigit   = false
		hasSpecial = false
	)

	// Check password length
	length := uint8(len(password))
	if length >= minLen {
		hasMinLen = true
	}
	if length <= maxLen {
		hasMaxLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return errors2.ErrPasswordMinLen
	}
	if !hasMaxLen {
		return errors2.ErrPasswordMaxLen
	}
	if !hasUpper {
		return errors2.ErrPasswordUpper
	}
	if !hasLower {
		return errors2.ErrPasswordLower
	}
	if !hasDigit {
		return errors2.ErrPasswordDigit
	}
	if !hasSpecial {
		return errors2.ErrPasswordSpecial
	}
	return nil
}
