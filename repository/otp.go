package repository

import (
	"reservation/config"
	"reservation/models"
	"reservation/pkg/utils"

	"github.com/google/uuid"
)

func CreateOtp(data models.Otp) (models.Otp, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetOtp(userID uuid.UUID, otpCode string, preloadFields []string) (response models.Otp, err error) {
	db := utils.BuildPreload(config.DB, preloadFields)

	err = db.Where("user_id = ? AND otp_code = ?", userID, otpCode).First(&response).Error

	return
}

func DeleteOtp(data models.Otp) error {
	err := config.DB.Delete(&data).Error
	return err
}
