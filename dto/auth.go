package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	NoHP     string `json:"noHp"`
	Password string `json:"password"`
}

func (request RegisterRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Fullname, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.NoHP, validation.Required),
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

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (request RefreshTokenRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.RefreshToken, validation.Required),
	)
}
