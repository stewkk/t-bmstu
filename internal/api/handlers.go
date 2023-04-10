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
	problem, err := handler.Service.GetProblem(problemId)
	if err != nil {
		return err
	}
	return ctx.JSONPretty(http.StatusOK, problem, identation)
}

func (handler *ServerInterfaceImpl) GetProblems(ctx echo.Context) error {
	return nil
}

func (handler *ServerInterfaceImpl) SumbitSolution(ctx echo.Context, problemId ProblemIdParameter) error {
	err := handler.Service.SumbitSolution(problemId)
	if err != nil {
		return err
	}
	return ctx.JSONPretty(http.StatusOK, "afsd", identation)
}

func (handler *ServerInterfaceImpl) GetSubmissions(ctx echo.Context) error {
	return nil
}

func (handler *ServerInterfaceImpl) GetSubmissionStatus(ctx echo.Context, submissionId SubmissionIdParameter) error {
	return nil
}

const identation = "    "
