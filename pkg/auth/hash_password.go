package auth

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// CreateHashPassword hashes the user's password for storage in the repository
func CreateHashPassword(password string) ([]byte, error) {
	logrus.Info("Creating Hash Password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("Failed to hash password")
		return nil, err
	}
	logrus.Info("Password hashed")
	return hashedPassword, nil
}

// ComparePasswords проверяет введенный пользователем пароль на соответствие сохраненному
func ComparePasswords(savedHashedPassword []byte, Password string) error {
	if err := bcrypt.CompareHashAndPassword(savedHashedPassword, []byte(Password)); err != nil {
		logrus.WithError(err).Error("Failed to compare hashed password")
		return err
	}
	return nil
}
