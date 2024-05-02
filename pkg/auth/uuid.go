package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// GenerateUUID generates a UUID using the github.com/google/uuid library
func GenerateUUID() uuid.UUID {
	return uuid.New()
}

// GetUUIDFromToken we check the validity of the token and if it is valid, then we get and return the UUID from it
func GetUUIDFromToken(tokenString string) (uuid.UUID, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signed method: %v", t.Header["alg"])
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		logrus.Error(err)
		return uuid.Nil, err
	}
	if !token.Valid {
		err = fmt.Errorf("token is not valid")
		logrus.Error(err)
		return uuid.Nil, err
	}
	logrus.Infof("Token is valid, userID: %v", claims.UserID)
	return claims.UserID, nil
}
