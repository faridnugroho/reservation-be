package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"reservation/dto"
	"reservation/models"
	"reservation/pkg/upload"
	"reservation/pkg/utils"
	"reservation/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UploadCarousel(file *multipart.FileHeader, userID string) (responseURL string, statusCode int, err error) {
	extension := filepath.Ext(file.Filename)
	if extension != ".png" && extension != ".jpg" && extension != ".jpeg" && extension != ".webp" {
		err = errors.New("the file extension is wrong. allowed file extensions are images (.png, .jpg, .jpeg, .webp)")
		statusCode = http.StatusBadRequest
		return
	}

	var src multipart.File
	src, err = file.Open()
	if err != nil {
		err = errors.New("faield to open file: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}
	defer src.Close()

	responseURL, err = upload.UploadFile(src, userID, "")
	if err != nil {
		statusCode = http.StatusInternalServerError
		return
	}

	data := models.Carousels{
		Url: responseURL,
	}

	_, err = repository.UploadCarousel(data)
	if err != nil {
		err = errors.New("failed to create data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}

func GetCarousels(param utils.PagingRequest, preloadFields []string) (response utils.PagingResponse, data []models.Carousels, statusCode int, err error) {
	baseFilter := "deleted_at IS NULL"
	filter := baseFilter
	var filterValues []any

	if param.Custom != "" {
		filter += " AND status = ?"
		filterValues = append(filterValues, param.Custom.(string))
	}
	if param.Search != "" {
		filter += " AND (url ILIKE ?)"
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
	}

	data, total, totalFiltered, err := repository.GetCarousels(dto.FindParameter{
		BaseFilter:   baseFilter,
		Filter:       filter,
		FilterValues: filterValues,
		Limit:        param.Limit,
		Order:        param.Order,
		Offset:       param.Offset,
	}, preloadFields)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	response = utils.PopulateResPaging(&param, data, total, totalFiltered)
	statusCode = http.StatusOK

	return
}

func DeleteCarousel(id string) (statusCode int, err error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		err = errors.New("failed to parse UUID: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	data, err := repository.GetCarouselByID(parsedUUID, []string{})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	err = repository.DeleteCarousel(data)
	if err != nil {
		err = errors.New("failed to delete data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}

func UpdateCarouselStatus(id string) (response models.Carousels, statusCode int, err error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		err = errors.New("failed to parse UUID: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	data, err := repository.GetCarouselByID(parsedUUID, []string{})
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if err == gorm.ErrRecordNotFound {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	data.Status = !data.Status

	response, err = repository.UpdateCarouselStatus(data)
	if err != nil {
		err = errors.New("failed to update data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}
