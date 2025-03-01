package service

import (
	"errors"
	"net/http"
	"reservation/dto"
	"reservation/models"
	"reservation/repository"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetOtp(request dto.VerifyEmailRequest, preloadFields []string) (data models.Otp, statusCode int, err error) {
	parsedUUID, err := uuid.Parse(request.UserID)
	if err != nil {
		err = errors.New("failed to parse UUID: " + err.Error())
		statusCode = http.StatusInternalServerError

		return
	}

	data, err = repository.GetOtp(parsedUUID, request.OtpCode, preloadFields)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			statusCode = http.StatusNotFound

			return
		}

		statusCode = http.StatusInternalServerError

		return
	}

	// Validasi OTP expiry (10 menit)
	expiryDuration := 10 * time.Minute
	if time.Since(data.CreatedAt) > expiryDuration {
		err = errors.New("otp has expired")
		statusCode = http.StatusBadRequest

		return
	}

	statusCode = http.StatusOK

	return
}
