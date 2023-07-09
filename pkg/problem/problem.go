package problem

import "errors"

type Problem struct {
	ProblemMeta
	Statement string
	Extra interface{}
}

type ProblemMeta struct {
	Id string
	Ts TestingSystemType
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

type TimusExtra struct {
	TimusId string
}

type CodeforcesExtra struct {
	CfId string
}
