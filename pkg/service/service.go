package service

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/stewkk/t-bmstu/internal/errors"
	"github.com/stewkk/t-bmstu/pkg/models"
	"github.com/stewkk/t-bmstu/pkg/timus"
)

type TestingSystemService struct {
	ts timus.TestingSystem
}

func (s *TestingSystemService) GetProblem(problemId models.ProblemId) (models.Problem, error) {
	idToLink := map[models.ProblemId]string {
		uuid.MustParse("7a81f316-add0-4dc9-85a7-28b8fd747f8f"):  "https://acm.timus.ru/problem.aspx?num=1000",
		uuid.MustParse("0960f297-8e9f-492e-90bd-8708772638e3"):  "https://acm.timus.ru/problem.aspx?num=1004",
		uuid.MustParse("e346caa7-5af4-4741-a1a0-d4f5f88a87c2"):  "https://acm.timus.ru/problem.aspx?num=1005",
	}
	url, ok := idToLink[problemId]
	if !ok {
		return models.Problem{}, errors.NewHTTPError(http.StatusNotFound, "ProblemId is not found in problems map")
	}
	return s.ts.GetProblem(url)
}
