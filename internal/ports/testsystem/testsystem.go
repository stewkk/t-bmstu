package testsystem

//go:generate mockery --name TestingSystem
type TestingSystem interface {
	SubmitSolution(*SubmissionCreateRequest) (*SubmissionCreateResponse, error)
}

type SubmissionCreateRequest struct {}

type SubmissionCreateResponse struct {
	SubmissionId string
}
