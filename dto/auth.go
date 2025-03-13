package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (request RegisterRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Fullname, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Phone, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
}

type LoginRequest struct {
	EmailOrPhone string `json:"emailOrPhone"`
	Password     string `json:"password"`
}

func (request LoginRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.EmailOrPhone, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
}

type ResendEmailVerificationRequest struct {
	UserID string `json:"userId"`
}

func (request ResendEmailVerificationRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.UserID, validation.Required),
	)
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (request RefreshTokenRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.RefreshToken, validation.Required),
	)
}

type SendForgotPasswordRequest struct {
	Email string `json:"email"`
}

func (request SendForgotPasswordRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required, is.Email),
	)
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}

func (request ResetPasswordRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.NewPassword, validation.Required),
	)
}
