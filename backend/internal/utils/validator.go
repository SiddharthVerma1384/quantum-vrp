package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// You can optionally format and customize the validation errors here
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}
