package repository

import (
	"reservation/config"
	"reservation/dto"
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

func GetCarousels(param dto.FindParameter, preloadsFields []string) (responses []models.Carousels, total int64, totalFiltered int64, err error) {
	err = config.DB.Model(responses).Where(param.BaseFilter, param.BaseFilterValues...).Count(&total).Error
	if err != nil {
		return
	}

	err = config.DB.Model(responses).Where(param.Filter, param.FilterValues...).Count(&totalFiltered).Error
	if err != nil {
		return
	}

	db := utils.BuildPreload(config.DB, preloadsFields)

	if param.Limit == 0 {
		err = db.Offset(param.Offset).Order(param.Order).Where(param.Filter, param.FilterValues...).Find(&responses).Error
	} else {
		err = db.Limit(param.Limit).Offset(param.Offset).Order(param.Order).Where(param.Filter, param.FilterValues...).Find(&responses).Error
	}

	return
}

func DeleteCarousel(data models.Carousels) error {
	err := config.DB.Delete(&data).Error

	return err
}

func UpdateCarouselStatus(data models.Carousels) (models.Carousels, error) {
	err := config.DB.Save(&data).Error

	return data, err
}
