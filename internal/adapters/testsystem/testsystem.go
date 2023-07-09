package testsystem

import (
	"github.com/stewkk/t-bmstu/internal/ports/testsystem"
	"github.com/stewkk/t-bmstu/pkg/problem"
)

// Make sure we conform to TestingSystem interface
// var _ testsystem.TestingSystem = (*TestingSystem)(nil)

func NewTestingSystem(problem problem.ProblemMeta) (testsystem.TestingSystem, error) {
	// TODO
	return nil, nil
}
