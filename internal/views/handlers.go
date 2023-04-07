package views

import (
	"io"
	"net/http"
	"html/template"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"

	"github.com/stewkk/t-bmstu/pkg/service"
)

type Template struct {
    Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

type ServerInterfaceImpl struct {
	Service service.TestingSystemService
}

func (handler *ServerInterfaceImpl) GetProblemView(ctx echo.Context, problemId openapi_types.UUID) error {
	problem, err := handler.Service.GetProblem(problemId)
	if err != nil {
		return err
	}
	return ctx.Render(http.StatusOK, "problem", problem)
}
