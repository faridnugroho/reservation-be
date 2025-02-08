package dto

import (
	"reservation/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CarouselRequest struct {
	Url    string `json:"url"`
	Status bool   `json:"status"`
}

func (request CarouselRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Url, validation.Required),
	)
}

type Carouselesponse struct {
	models.CustomPublicModel
	Url    string `json:"url"`
	Status bool   `json:"status"`
}
