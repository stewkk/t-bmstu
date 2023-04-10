package views

import (
	"html/template"
	"io"
	"net/http"

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

func (handler *ServerInterfaceImpl) GetSubmissionsView(ctx echo.Context, problemId ProblemIdParameter) error {
	submissions, err := handler.Service.GetSubmissions()
	if err != nil {
		return err
	}
	return ctx.Render(http.StatusOK, "submissions", submissions)
}

func (handler *ServerInterfaceImpl) GetProblemsView(ctx echo.Context) error {
	problems, err := handler.Service.GetProblems()
	if err != nil {
		return err
	}
	return ctx.Render(http.StatusOK, "problems", problems)
}
