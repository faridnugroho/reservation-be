package jwt

import (
	"errors"
	"fmt"
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

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		err = errors.New("failed to verify token: " + err.Error())
		return token, err
	}
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)

	claims, isOk := token.Claims.(jwt.MapClaims)
	if err != nil || !isOk || !token.Valid {
		err = errors.New("failed to decode token: " + err.Error())
		return claims, err
	}

	return claims, nil
}
