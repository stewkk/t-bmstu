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

	languages := []string{"FreePascal 2.6",
		"Visual C 2019",
		"Visual C++ 2019",
		"Visual C 2019 x64",
		"Visual C++ 2019 x64",
		"GCC 9.2 x64",
		"G++ 9.2 x64",
		"Clang++ 10 x64",
		"Java 1.8",
		"Visual C# 2019",
		"Python 3.8 x64",
		"PyPy 3.8 x64",
		"Go 1.14 x64",
		"Ruby 1.9",
		"Haskell 7.6",
		"Scala 2.11",
		"Rust 1.58 x64",
		"Kotlin 1.4.0",
	}

	return ctx.Render(http.StatusOK, "problem", Problem{
		Id:        problemId,
		Name:      "TODO",
		Statement: template.HTML(problem.Statement),
		Languages: languages,
	})
}

func (s *Server) GetSubmissionsView(ctx echo.Context, problemId ProblemIdParameter) error {
	return nil
}
