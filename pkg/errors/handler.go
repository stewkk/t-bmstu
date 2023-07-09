package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, ctx echo.Context) {
	res := New(http.StatusInternalServerError, "Unexpected error").(*Error)
	if he, ok := err.(*echo.HTTPError); ok {
		res.Status = he.Code
		res.Title = he.Message.(string)
	}
	if he, ok := err.(*Error); ok {
		res.Status = he.Status
		res.Title = he.Title
	}
	ctx.Logger().Errorf("%v: %v", res.Status, res.Title)

	tag := ctx.Get("tag")
	if v, ok := tag.(string); ok && v == "api" {
		ctx.JSONPretty(res.Status, res, "    ")
		return
	}
	ctx.Render(res.Status, "error", res)
}
