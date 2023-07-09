package rest

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/stewkk/t-bmstu/internal/app"
	"github.com/stewkk/t-bmstu/pkg/errors"
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
	var input SubmitBody
	err := ctx.Bind(&input)
	if err != nil {
		return errors.New(http.StatusBadRequest, err.Error())
	}

	meta, err := s.App.SubmitSolution(app.SubmissionInput{
		ProblemId: problemId.String(),
		Body:      input.SourceCode,
		Language:  input.Language,
	})
	if err != nil {
		return err
	}

	// TODO: add some way to configure base address
	link := fmt.Sprintf("%v/api/submissions/%v", "https://t-bmstu.starovoytovai.ru", meta.Id)
	return ctx.JSONPretty(http.StatusCreated, CreatedSubmission{
		Id:   uuid.MustParse(meta.Id),
		Link: &link,
	}, identation)
}

func (s *Server) GetSubmissions(ctx echo.Context) error {
	return nil
}

func (s *Server) GetSubmissionStatus(ctx echo.Context, submissionId SubmissionIdParameter) error {
	return nil
}

// 4 spaces
const identation = "    "
