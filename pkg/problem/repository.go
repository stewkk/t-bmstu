package problem

//go:generate mockery --name Repository
type Repository interface {
	GetProblems() ([]ProblemMeta, error)
	GetProblem(string) (Problem, error)
}
