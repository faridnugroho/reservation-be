package service

import (
	"errors"
	"fmt"
	"log"
	"reservation/config"
	"reservation/dto"
	"reservation/models"
	webToken "reservation/pkg/jwt"
	"reservation/pkg/smtp"
	"reservation/pkg/utils"
	"reservation/repository"
	"strings"

	"github.com/google/uuid"
)

func SendEmailVerification(userID uuid.UUID, targetEmail string) error {
	subject := "OTP Register Reservation App"
	otpCode, err := utils.GenerateOTP()
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	data := models.Otp{
		OtpCode: otpCode,
		UserID:  userID,
	}

	_, err = repository.CreateOtp(data)
	if err != nil {
		return fmt.Errorf("failed to create data: %w", err)
	}

	go smtp.SendEmail("", targetEmail, subject, otpCode)

	return nil
}

func VerifyUser(request dto.VerifyEmailRequest) (dataUser models.Users, statusCode int, err error) {
	dataOtp, statusCode, err := GetOtp(request, []string{})
	if err != nil {
		log.Println(err.Error())

		return
	}

	dataUser, _, err = GetUserByID(request.UserID, []string{})
	if err != nil {
		log.Println(err.Error())

		return
	}

	dataUser.IsVerified = true
	dataUser, err = repository.UpdateUser(dataUser)
	if err != nil {
		log.Println("Failed to update user: " + err.Error())

		return
	}

	err = repository.DeleteOtp(dataOtp)
	if err != nil {
		log.Println("Failed to delete OTP: " + err.Error())
	}

	return
}

func RefreshToken(expiredAccessToken string) (string, error) {
	if expiredAccessToken == "" {
		return "", errors.New("no JWT token provided")
	}

	claims, err := webToken.DecodeToken(expiredAccessToken)
	if err != nil {
		if !strings.Contains(err.Error(), "expired") {
			return "", errors.New("failed to decode token: " + err.Error())
		}
	}

	claims["exp"] = config.LoadConfig().JWTExpirationTime

	newAccessToken, err := webToken.GenerateToken(&claims)
	if err != nil {
		return "", errors.New("failed to generate token: " + err.Error())
	}

	return newAccessToken, nil
}
