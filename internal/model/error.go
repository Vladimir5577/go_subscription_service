package model

type ValidationErrorResponse struct {
	ValidationError []string `json:"validation_error"`
}
