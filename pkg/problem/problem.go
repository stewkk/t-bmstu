package problem

import "errors"

type Problem struct {
	ProblemMeta
	Statement string
}

type ProblemMeta struct {
	Id string
	Ts TestingSystemType
	ExternalId string
	ExternalLink string
}

//go:generate stringer -type=TestingSystemType
type TestingSystemType int

const (
	Timus TestingSystemType = iota
)

func NewTestingSystemType(s string) (TestingSystemType, error) {
	for i, e := range [1]string{"Timus"} {
		if s == e {
			return TestingSystemType(i), nil
		}
	}
	return -1, errors.New("bad TestingSystemType")
}
