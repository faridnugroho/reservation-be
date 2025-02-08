package service

import (
	"errors"
	"reservation/config"
	"reservation/pkg/jwt"
	"strings"
)

func RefreshToken(expiredAccessToken string) (string, error) {
	if expiredAccessToken == "" {
		return "", errors.New("no JWT token provided")
	}

	claims, err := jwt.DecodeToken(expiredAccessToken)
	if err != nil {
		if !strings.Contains(err.Error(), "expired") {
			return "", errors.New("failed to decode token: " + err.Error())
		}
	}

	claims["exp"] = config.LoadConfig().JWTExpirationTime

	newAccessToken, err := jwt.GenerateToken(&claims)
	if err != nil {
		return "", errors.New("failed to generate token: " + err.Error())
	}

	return newAccessToken, nil
}
