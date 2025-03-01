package service

import (
	"errors"
	"fmt"
	"net/http"
	"reservation/dto"
	"reservation/models"
	"reservation/pkg/bcrypt"
	"reservation/pkg/utils"
	"reservation/repository"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(request dto.UserRequest) (response models.Users, statusCode int, err error) {
	data := models.Users{
		Fullname: request.Fullname,
		Email:    request.Email,
		No_hp:    request.NoHP,
		Password: bcrypt.HashPassword(request.Password),
	}

	response, err = repository.CreateUser(data)
	if err != nil {
		err = errors.New("failed to create data: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusCreated
	return
}

func GetUsers(fullname, email string, param utils.PagingRequest, preloadFields []string) (response utils.PagingResponse, data []models.Users, statusCode int, err error) {
	baseFilter := "deleted_at IS NULL"
	filter := baseFilter
	var filterValues []any

	if fullname != "" {
		filter += " AND fullname = ?"
		filterValues = append(filterValues, fullname)
	}
	if email != "" {
		filter += " AND (email = ? OR no_hp = ?)"
		filterValues = append(filterValues, email, email)
	}
	if param.Search != "" {
		filter += " AND (fullname ILIKE ? OR email ILIKE ? OR no_hp ILIKE ?)"
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
		filterValues = append(filterValues, fmt.Sprintf("%%%s%%", param.Search))
	}

	data, total, totalFiltered, err := repository.GetUsers(dto.FindParameter{
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

func GetUserByID(id string, preloadFields []string) (data models.Users, statusCode int, err error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		err = errors.New("failed to parse UUID: " + err.Error())
		statusCode = http.StatusInternalServerError
		return
	}
	data, err = repository.GetUserByID(parsedUUID, preloadFields)
	if err != nil {
		err = errors.New("failed to get data: " + err.Error())
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			statusCode = http.StatusNotFound
			return
		}

		statusCode = http.StatusInternalServerError
		return
	}

	statusCode = http.StatusOK
	return
}
