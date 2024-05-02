package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

// Claims — claims structure that includes standard claims and UserID
type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

const (
	TokenExp  = time.Hour * 3
	SecretKey = "SnJSkf123jlLKNfsNln"
)

// BuildJWTString creates a token with the HS256 signature algorithm and Claims statements and returns it as a string.
func BuildJWTString(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// time create token
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: userID,
	})
	// create token string
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return tokenString, nil
}
