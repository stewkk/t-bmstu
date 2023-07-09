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
	tid, err := getTimusId(id)
	if err != nil {
		return problem.Problem{}, err
	}
	link := "https://acm.timus.ru/problem.aspx?num="+tid

	ts := timus.TestingSystem{}
	p, err := ts.GetProblem(link)
	return problem.Problem{
		ProblemMeta: problem.ProblemMeta{
			Id:           id,
			Ts:           problem.Timus,
			ExternalId:   tid,
			ExternalLink: link,
		},
		Statement:   p.Statement,
	}, err
}

var idToTimusId = map[string]string{
	uuid.MustParse("7a81f316-add0-4dc9-85a7-28b8fd747f8f").String(): "1000",
	uuid.MustParse("0960f297-8e9f-492e-90bd-8708772638e3").String(): "1004",
	uuid.MustParse("e346caa7-5af4-4741-a1a0-d4f5f88a87c2").String(): "1005",
}

func getTimusId(id string) (string, error) {
	timusId, ok := idToTimusId[id]
	if !ok {
		return "", errors.New(http.StatusNotFound, "ProblemId is not found")
	}
	return timusId, nil
}
