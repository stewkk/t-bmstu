package service

import (
	"github.com/stewkk/t-bmstu/pkg/models"
	"github.com/stewkk/t-bmstu/pkg/timus"
)

type TestingSystemService struct {}

func (s *TestingSystemService) GetProblem(problemId models.ProblemId) models.Problem {
	impl := s.getImpl(problemId)
	return impl.GetProblem(problemId)
}

func (s *TestingSystemService) getImpl(problemId models.ProblemId) TestingSystemImpl {
	return &timus.TestingSystem{}
}
