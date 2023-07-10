package timus

import (
	"github.com/google/uuid"

	"github.com/stewkk/t-bmstu/internal/ports/testsystem"
	"github.com/stewkk/t-bmstu/pkg/testsystems/timus"
)

// Make sure we conform to TestingSystem interface
var _ testsystem.TestingSystem = (*TestingSystem)(nil)

type TestingSystem struct{}

// SubmitSolution implements testsystem.TestingSystem
func (*TestingSystem) SubmitSolution(r *testsystem.SubmissionCreateRequest) (*testsystem.SubmissionCreateResponse, error) {
	ts := timus.TestingSystem{}
	// TODO: configure judje id somehow
	err := ts.SendSubmission("342187EL", toTimusLang(r.Language), r.ExternalId, r.SourceCode)
	if err != nil {
		return nil, err
	}
	return &testsystem.SubmissionCreateResponse{
		SubmissionId: uuid.NewString(),
	}, nil
}

func toTimusLang(l string) string {
	n := map[string]string{
		"FreePascal 2.6":      "31",
		"Visual C 2019":       "63",
		"Visual C++ 2019":     "64",
		"Visual C 2019 x64":   "65",
		"Visual C++ 2019 x64": "66",
		"GCC 9.2 x64":         "67",
		"G++ 9.2 x64":         "68",
		"Clang++ 10 x64":      "69",
		"Java 1.8":            "32",
		"Visual C# 2019":      "61",
		"Python 3.8 x64":      "57",
		"PyPy 3.8 x64":        "71",
		"Go 1.14 x64":         "58",
		"Ruby 1.9":            "18",
		"Haskell 7.6":         "19",
		"Scala 2.11":          "33",
		"Rust 1.58 x64":       "72",
		"Kotlin 1.4.0":        "60",
	}

	return n[l]
}
