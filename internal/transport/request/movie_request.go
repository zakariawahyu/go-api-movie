package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateMovieRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
}

func (request CreateMovieRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.Description, validation.Required),
		validation.Field(&request.Rating, validation.Required),
		validation.Field(&request.Image, validation.Required),
	)
}

type UpdateMovieRequest struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
}

func (request UpdateMovieRequest) Validate() error {
	return validation.ValidateStruct(
		&request,
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.Description, validation.Required),
		validation.Field(&request.Rating, validation.Required),
		validation.Field(&request.Image, validation.Required),
	)
}
