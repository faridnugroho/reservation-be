package controllers

import (
	"net/http"
	"reservation/dto"
	"reservation/pkg/utils"
	"reservation/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func UploadCarousel(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(
			http.StatusUnauthorized,
			dto.Response{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized: currentUser not found",
				Error:   "User authentication failed",
			},
		)

		return
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Invalid token claims format",
				Error:   "Failed to parse user claims",
			},
		)

		return
	}

	userID, ok := claims["id"].(string)
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "User ID is not a string",
				Error:   "Invalid user ID format",
			},
		)

		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Failed to get file from form",
				Error:   err.Error(),
			},
		)

		return
	}

	responseURL, statusCode, err := service.UploadCarousel(file, userID)
	if err != nil {
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to upload product picture",
				Error:   err.Error(),
			},
		)

		return
	}

	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to upload",
			Data:    responseURL,
		},
	)
}

func GetCarousels(c *gin.Context) {
	var preloadFields []string
	param := utils.PopulatePaging(c, "status")

	data, _, statusCode, err := service.GetCarousels(param, preloadFields)

	if err != nil {
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to get data",
				Error:   err.Error(),
			},
		)

		return
	}

	c.JSON(statusCode, data)
}

func UpdateCarousel(c *gin.Context) {
	id := c.Param("id")

	var request dto.CarouselRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Error:   err.Error(),
			})

		return
	}

	data, statusCode, err := service.UpdateCarousel(id, request)
	if err != nil {
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to update data",
				Error:   err.Error(),
			},
		)

		return
	}

	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to update data",
			Data:    data,
		},
	)
}

func DeleteCarousel(c *gin.Context) {
	id := c.Param("id")

	statusCode, err := service.DeleteCarousel(id)
	if err != nil {
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to delete data",
				Error:   err.Error(),
			},
		)

		return
	}

	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to delete data",
		},
	)
}

func UpdateCarouselStatus(c *gin.Context) {
	id := c.Param("id")

	var request dto.CarouselRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request body",
				Error:   err.Error(),
			})

		return
	}

	data, statusCode, err := service.UpdateCarouselStatus(id)
	if err != nil {
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to update data",
				Error:   err.Error(),
			},
		)

		return
	}

	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to update data",
			Data:    data,
		},
	)
}
