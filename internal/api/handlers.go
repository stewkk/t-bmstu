package api

import (
	"net/http"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
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
	return ctx.JSONPretty(http.StatusOK, problem, "    ")
}
