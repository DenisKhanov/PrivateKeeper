package validate

import (
	"github.com/DenisKhanov/PrivateKeeper/internal/domain"
	"github.com/sirupsen/logrus"
)

// CheckRegistrationData combines CheckLogin, CheckPassword, and CheckEmail;
// if the login, password, or email do not comply, it returns the appropriate errors
func CheckRegistrationData(login string, password string, email string) error {
	if err := CheckLogin(login, domain.MinLoginLength, domain.MaxLoginLength); err != nil {
		logrus.Error(err)
		return err
	}
	if err := CheckPassword(password, domain.MinPasswordLength, domain.MaxPasswordLength); err != nil {
		logrus.Error(err)
		return err
	}
	if err := CheckEmail(email); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
