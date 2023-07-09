package testsystem

//go:generate mockery --name TestingSystem
type TestingSystem interface {
	SubmitSolution(*SubmissionCreateRequest) (*SubmissionCreateResponse, error)
}

type SubmissionCreateRequest struct {
	ProblemId string
	SourceCode string
	Language string
	Extra interface{}
}

type SubmissionCreateResponse struct {
	SubmissionId string
}
