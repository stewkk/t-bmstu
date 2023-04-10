package html

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

func (s *Server) GetProblemsView(ctx echo.Context) error {
	return nil
}

func (s *Server) GetProblemView(ctx echo.Context, problemId ProblemIdParameter) error {
	problem, err := s.App.GetProblem(problemId.String())
	if err != nil {
		return err
	}

	return ctx.Render(http.StatusOK, "problem", Problem{
		Id:        problemId,
		Name:      "TODO",
		Statement: template.HTML(problem.Statement),
	})
}

func (s *Server) GetSubmissionsView(ctx echo.Context, problemId ProblemIdParameter) error {
	return nil
}
