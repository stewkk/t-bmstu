package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/stewkk/t-bmstu/pkg/models"
)

func ErrorHandler(err error, ctx echo.Context) {
	code := http.StatusInternalServerError
	message := "Unexpected error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}
	ctx.Logger().Error(err)

	error := models.Error{
		Code:    code,
		Message: message,
	}

	tag := ctx.Get("tag")
	if v, ok := tag.(string); ok && v == "api" {
		ctx.JSONPretty(code, error, "    ")
		return
	}
	ctx.Render(code, "error", error)
}
