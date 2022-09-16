package utils

import (
	"net/http"
	"ofa/bin/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidationUtil struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return response.ReplyError(err.(*echo.HTTPError).Internal.Error(), http.StatusBadRequest)
	}
	if err := c.Validate(i); err != nil {
		return response.ReplyError(err.(validator.ValidationErrors)[0].Error(), http.StatusBadRequest)
	}
	
	return nil
}

