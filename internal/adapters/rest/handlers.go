package rest

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stewkk/t-bmstu/internal/app"
)

// Make sure we conform to ServerInterface interface
var _ ServerInterface = (*Server)(nil)

type Server struct {
	App app.App
}

func (s *Server) Ping(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (s *Server) GetProblems(ctx echo.Context) error {
	return nil
}

func (s *Server) GetProblem(ctx echo.Context, problemId ProblemIdParameter) error {
	problem, err := s.App.GetProblem(problemId.String())
	if err != nil {
		return err
	}

	return ctx.JSONPretty(http.StatusOK, Problem{
		Id:        problemId,
		Name:      "TODO",
		Statement: template.HTML(problem.Statement),
	}, identation)
}

func (s *Server) SumbitSolution(ctx echo.Context, problemId ProblemIdParameter) error {
	return nil
}

func (s *Server) GetSubmissions(ctx echo.Context) error {
	return nil
}

func (s *Server) GetSubmissionStatus(ctx echo.Context, submissionId SubmissionIdParameter) error {
	return nil
}

// 4 spaces
const identation = "    "
