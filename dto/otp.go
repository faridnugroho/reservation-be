package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type VerifyEmailRequest struct {
	UserID  string `json:"userId"`
	OtpCode string `json:"otpCode"`
}

func (request VerifyEmailRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.UserID, validation.Required),
		validation.Field(&request.OtpCode, validation.Required),
	)
}
