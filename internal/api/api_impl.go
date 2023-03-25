package api

import (
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"

	"github.com/stewkk/t-bmstu/pkg/service"
)

type ServerInterfaceImpl struct {
	Service service.TestingSystemService
}

func (handler *ServerInterfaceImpl) Ping(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (handler *ServerInterfaceImpl) GetProblem(ctx echo.Context, problemId openapi_types.UUID) error {
	problem := handler.Service.GetProblem(problemId)
	return handler.renderWithView(ctx, http.StatusOK, problem, "problem")
}

func (handler *ServerInterfaceImpl) renderWithView(ctx echo.Context, code int, body interface{}, template string) error {
	if ctx.Request().Header.Get("Accept") == "text/html" {
		return ctx.Render(code, template, body)
	}
	return handler.render(ctx, code, body)
}

func (handler *ServerInterfaceImpl) render(ctx echo.Context, code int, body interface{}) error {
	yamlHeader := "application/x-yaml"
	switch ctx.Request().Header.Get("Accept") {
	case yamlHeader:
		blob, err := yaml.Marshal(body)
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.Blob(code, yamlHeader, blob)
	case "application/json":
		return ctx.JSONPretty(code, body, "    ")
	}
	return ctx.NoContent(http.StatusNotAcceptable)
}
