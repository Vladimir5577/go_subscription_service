package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"subscriptions_service/internal/model"

	"github.com/go-playground/validator/v10"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		JsonResponse(*w, err.Error(), 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		errorResponse := ParseValidationErrors(err)
		JsonResponse(*w, errorResponse, 402)
		return nil, err
	}

	return &body, nil
}

func ParseValidationErrors(err error) model.ValidationErrorResponse {
	var errs []string
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Printf("%#v", err)
		errString := fmt.Errorf("field %s: must be %s %s, got `%s`", err.Field(), err.Tag(), err.Param(), err.Value())
		errs = append(errs, errString.Error())
	}
	errorResponse := model.ValidationErrorResponse{
		ValidationError: errs,
	}
	return errorResponse
}
