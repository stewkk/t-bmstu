package errors

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func NewHTTPError(code int, message string) error {
	return echo.NewHTTPError(code, message)
}

func NewError(message string) error {
	return errors.New(message)
}
