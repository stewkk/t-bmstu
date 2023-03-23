package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

type ServerInterfaceImpl struct {}



func (handler *ServerInterfaceImpl) Ping(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (handler *ServerInterfaceImpl) GetProblem(ctx echo.Context, problemId openapi_types.UUID) error {
	return nil
}

func (handler *ServerInterfaceImpl) GetProblemView(ctx echo.Context, problemId openapi_types.UUID) error {
	return nil
}
