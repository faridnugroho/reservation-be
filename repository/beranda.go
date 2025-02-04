package repository

import (
	"reservation/config"
	"reservation/models"
	"reservation/pkg/utils"

	"github.com/google/uuid"
)

func UploadCarousel(data models.Carousels) (models.Carousels, error) {
	err := config.DB.Create(&data).Error

	return data, err
}

func GetCarouselByID(id uuid.UUID, preloadFields []string) (response models.Carousels, err error) {
	db := utils.BuildPreload(config.DB, preloadFields)

	err = db.Where("id = ?", id).First(&response).Error

	return
}

func UpdateCarouselStatus(data models.Carousels) (models.Carousels, error) {
	err := config.DB.Save(&data).Error

	return data, err
}
