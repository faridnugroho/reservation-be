package controllers

import (
	"net/http"
	"reservation/config"
	"reservation/dto"
	"reservation/pkg/bcrypt"
	webToken "reservation/pkg/jwt"
	"reservation/pkg/utils"
	"reservation/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *gin.Context) {
	var request dto.RegisterRequest

	// Bind the request
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			dto.Response{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid request body",
				Error:   err.Error(),
			},
		)

		return
	}

	// Validate the request
	if err := request.Validate(); err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request value",
				Error:   err.Error(),
			},
		)

		return
	}

	// Check if email has been registered
	param := utils.PopulatePaging(c, "")
	_, check, _, _ := service.GetUsers("", request.Email, param, []string{})
	if len(check) > 0 {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Email has been registered",
				Error:   "",
			},
		)

		return
	}

	// Check if phone number has been registered
	_, checkPhone, _, _ := service.GetUsers("", request.NoHP, param, []string{})
	if len(checkPhone) > 0 {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Phone number has been registered",
				Error:   "",
			},
		)
		return
	}

	// Call the service layer
	result, statusCode, err := service.CreateUser(dto.UserRequest(request))
	if err != nil {
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to register",
				Error:   err.Error(),
			},
		)

		return
	}

	// send email verification
	go service.SendEmailVerification(result.ID, request.Email)

	// Return the response
	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to register",
			Data:    result,
		},
	)
}

func Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			dto.Response{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid request body",
				Error:   err.Error(),
			},
		)

		return
	}

	if err := request.Validate(); err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request value",
				Error:   err.Error(),
			},
		)

		return
	}

	param := utils.PopulatePaging(c, "")
	_, user, statusCode, _ := service.GetUsers("", request.EmailOrPhone, param, []string{})
	if len(user) == 0 {
		c.JSON(
			http.StatusNotFound,
			dto.Response{
				Status:  http.StatusNotFound,
				Message: "Email not found",
			},
		)

		return
	}

	err := bcrypt.VerifyPassword(request.Password, user[0].Password)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Failed to verify password",
				Error:   err.Error(),
			},
		)

		return
	}

	claims := jwt.MapClaims{}
	claims["id"] = user[0].ID
	claims["exp"] = config.LoadConfig().JWTExpirationTime

	token, err := webToken.GenerateToken(&claims)
	if err != nil {
		c.JSON(
			http.StatusUnauthorized,
			dto.Response{
				Status:  401,
				Message: "Failed to generate jwt token",
				Error:   err.Error(),
			},
		)

		return
	}

	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to login",
			Data:    token,
		},
	)
}

func VerifyUser(c *gin.Context) {
	var request dto.VerifyEmailRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			dto.Response{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid request body",
				Error:   err.Error(),
			},
		)

		return
	}

	if err := request.Validate(); err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request value",
				Error:   err.Error(),
			},
		)

		return
	}

	otp, statusCode, err := service.VerifyUser(dto.VerifyEmailRequest(request))
	if err != nil {
		c.JSON(
			statusCode,
			dto.Response{
				Status:  statusCode,
				Message: "Failed to verify email",
				Error:   err.Error(),
			},
		)

		return
	}

	c.JSON(
		statusCode,
		dto.Response{
			Status:  statusCode,
			Message: "Success to verify email",
			Data:    otp,
		},
	)
}

func RefreshToken(c *gin.Context) {
	var request dto.RefreshTokenRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(
			http.StatusUnprocessableEntity,
			dto.Response{
				Status:  http.StatusUnprocessableEntity,
				Message: "Invalid request body",
				Error:   err.Error(),
			},
		)

		return
	}

	if err := request.Validate(); err != nil {
		c.JSON(
			http.StatusBadRequest,
			dto.Response{
				Status:  http.StatusBadRequest,
				Message: "Invalid request value",
				Error:   err.Error(),
			},
		)

		return
	}

	newAccessToken, err := service.RefreshToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Kirimkan token baru
	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
