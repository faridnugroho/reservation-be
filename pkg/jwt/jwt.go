package jwt

import (
	"errors"
	"reservation/config"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = config.LoadConfig().SecretKey

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		err = errors.New("failed to generate token: " + err.Error())
		return "", err
	}

	return webtoken, nil
}
