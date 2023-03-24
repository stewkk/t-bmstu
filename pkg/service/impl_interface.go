package service

import "github.com/stewkk/t-bmstu/pkg/models"

type TestingSystemImpl interface {
	GetProblem(problemId models.ProblemId) models.Problem
}
