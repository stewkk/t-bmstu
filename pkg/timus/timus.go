package timus

import (
	"github.com/stewkk/t-bmstu/pkg/models"
)

type TestingSystem struct {}

func (ts *TestingSystem) GetProblem(problemId models.ProblemId) models.Problem {
	return models.Problem{
		Statement: "TODO",
	}
}
