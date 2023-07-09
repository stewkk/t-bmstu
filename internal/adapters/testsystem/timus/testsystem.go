package timus

import (
	"github.com/google/uuid"

	"github.com/stewkk/t-bmstu/internal/ports/testsystem"
	"github.com/stewkk/t-bmstu/pkg/problem"
	"github.com/stewkk/t-bmstu/pkg/testsystems/timus"
)

// Make sure we conform to TestingSystem interface
var _ testsystem.TestingSystem = (*TestingSystem)(nil)

type TestingSystem struct{}

// SubmitSolution implements testsystem.TestingSystem
func (*TestingSystem) SubmitSolution(r *testsystem.SubmissionCreateRequest) (*testsystem.SubmissionCreateResponse, error) {
	extra := r.Extra.(problem.TimusExtra)

	ts := timus.TestingSystem{}
	// TODO: configure judje id somehow
	err := ts.SendSubmission("342187EL", toTimusLang(r.Language), extra.TimusId, r.SourceCode)
	if err != nil {
		return nil, err
	}
	return &testsystem.SubmissionCreateResponse{
		SubmissionId: uuid.NewString(),
	}, nil
}

func toTimusLang(l string) string {
	// TODO: only python for now
	return "57"
}
