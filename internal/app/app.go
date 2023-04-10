package app

import (
	"github.com/stewkk/t-bmstu/internal/ports/testsystem"
	"github.com/stewkk/t-bmstu/pkg/problem"
)

type App struct {
	problemRepo problem.Repository
	ts testsystem.TestingSystem
}

func NewApp(problemRepo problem.Repository, ts testsystem.TestingSystem) App {
	return App{
		problemRepo: problemRepo,
		ts:          ts,
	}
}

func (a *App) GetProblems() ([]ProblemMetaDto, error) {
	problems, err := a.problemRepo.GetProblems()
	if err != nil {
		return nil, err
	}

	problemsDto := []ProblemMetaDto{}
	for _, problem := range problems {
		problemsDto = append(problemsDto, ToProblemMetaDto(problem))
	}
	return problemsDto, nil
}

func (a* App) GetProblem(id string) (ProblemDto, error) {
	p, err := a.problemRepo.GetProblem(id)
	if err != nil {
		return ProblemDto{}, err
	}

	return ToProblemDto(p), nil
}

func (a *App) SubmitSolution(input SubmissionInput) (SubmissionMetaDto, error) {
	r, err := a.ts.SubmitSolution(&testsystem.SubmissionCreateRequest{})
	if err != nil {
		return SubmissionMetaDto{}, err
	}

	return SubmissionMetaDto{
		Id: r.SubmissionId,
	}, nil
}
