package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	No_hp    string `json:"no_hp"`
	Password string `json:"password"`
}

func (request RegisterRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Fullname, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.No_hp, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request LoginRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Required),
	)
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (request RefreshTokenRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.RefreshToken, validation.Required),
	)
}
