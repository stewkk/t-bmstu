package app

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/stewkk/t-bmstu/internal/ports/testsystem"
	"github.com/stewkk/t-bmstu/pkg/errors"
	"github.com/stewkk/t-bmstu/pkg/problem"

	tsmock "github.com/stewkk/t-bmstu/internal/ports/testsystem/mocks"
	problemmock "github.com/stewkk/t-bmstu/pkg/problem/mocks"
)

type AppSuite struct {
	suite.Suite
	problemRepo *problemmock.Repository
	ts *tsmock.TestingSystem
	app App
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppSuite))
}

func (s *AppSuite) SetupTest() {
	s.ts = tsmock.NewTestingSystem(s.T())
	s.problemRepo = problemmock.NewRepository(s.T())
	s.app = NewApp(s.problemRepo, s.ts)
}

func (s *AppSuite) TestGetProblemsReturnsProblems() {
	s.problemRepo.EXPECT().GetProblems().Return([]problem.ProblemMeta{{}}, nil)

	got, _ := s.app.GetProblems()

	assert.Equal(s.T(), []ProblemMetaDto{{}}, got)
}

func (s *AppSuite) TestGetProblemHandlesDbError() {
	want := errors.New(http.StatusInternalServerError, "Unexpected error")
	s.problemRepo.EXPECT().GetProblems().Return(nil, want)
	var got *errors.Error

	_, err := s.app.GetProblems()

	assert.ErrorAs(s.T(), err, &got)
	assert.Equal(s.T(), want, got)
}

func (s *AppSuite) TestGetProblemReturnsProblem() {
	s.problemRepo.EXPECT().GetProblem(mock.Anything).Return(problem.Problem{
		ProblemMeta: problem.ProblemMeta{
			Id: "some-id",
		},
	}, nil)

	got, _ := s.app.GetProblem("some-id")

	assert.Equal(s.T(), ProblemDto{
		ProblemMetaDto: ProblemMetaDto{
			Id: "some-id",
		},
	}, got)
}

func (s *AppSuite) TestGetProblemHandlesNotFound() {
	want := errors.New(http.StatusNotFound, "Problem some-id not found")
	s.problemRepo.EXPECT().GetProblem(mock.Anything).Return(problem.Problem{}, want)
	var got *errors.Error

	_, err := s.app.GetProblem("some-id")

	assert.ErrorAs(s.T(), err, &got)
	assert.Equal(s.T(), want, got)
}

func (s *AppSuite) TestSumbitSolutionReturnsSubmissionWithId() {
	s.ts.EXPECT().SubmitSolution(mock.Anything).Return(&testsystem.SubmissionCreateResponse{
		SubmissionId: "some-id",
	}, nil).Once()
	s.problemRepo.EXPECT().GetProblem("problem-id").Return(problem.Problem{}, nil)

	got, _ := s.app.SubmitSolution(SubmissionInput{
		ProblemId: "problem-id",
	})

	assert.Equal(s.T(), SubmissionMetaDto{
		Id: "some-id",
	}, got)
}

func (s *AppSuite) TestSubmitSolutionHandlesTestSystemError() {
	want := errors.New(http.StatusNotFound, "Problem with id non-existing-id not found")
	s.problemRepo.EXPECT().GetProblem("non-existing-id").Return(problem.Problem{}, want)
	var got *errors.Error

	_, err := s.app.SubmitSolution(SubmissionInput{
		ProblemId: "non-existing-id",
	})

	assert.ErrorAs(s.T(), err, &got)
	assert.Equal(s.T(), want, got)
}
