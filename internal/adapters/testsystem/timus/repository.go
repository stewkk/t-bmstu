package timus

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/stewkk/t-bmstu/pkg/errors"
	"github.com/stewkk/t-bmstu/pkg/problem"
	"github.com/stewkk/t-bmstu/pkg/testsystems/timus"
)

// Make sure we conform to TestingSystem interface
var _ problem.Repository = (*WebProblemRepo)(nil)

type WebProblemRepo struct {}

func (r *WebProblemRepo) GetProblems() ([]problem.ProblemMeta, error) {
	return nil, nil
}

func (r *WebProblemRepo) GetProblem(id string) (problem.Problem, error) {
	url, err := getUrl(id)
	if err != nil {
		return problem.Problem{}, err
	}

	ts := timus.TestingSystem{}
	p, err := ts.GetProblem(url)
	return problem.Problem{
		ProblemMeta: problem.ProblemMeta{Id: id, Ts: problem.Timus},
		Statement:   p.Statement,
	}, err
}

func getUrl(id string) (string, error) {
	idToLink := map[string]string{
		uuid.MustParse("7a81f316-add0-4dc9-85a7-28b8fd747f8f").String(): "https://acm.timus.ru/problem.aspx?num=1000",
		uuid.MustParse("0960f297-8e9f-492e-90bd-8708772638e3").String(): "https://acm.timus.ru/problem.aspx?num=1004",
		uuid.MustParse("e346caa7-5af4-4741-a1a0-d4f5f88a87c2").String(): "https://acm.timus.ru/problem.aspx?num=1005",
	}
	url, ok := idToLink[id]
	if !ok {
		return "", errors.New(http.StatusNotFound, "ProblemId is not found")
	}
	return url, nil
}
