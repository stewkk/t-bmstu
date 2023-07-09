package app

import (
	"github.com/stewkk/t-bmstu/pkg/problem"
	"github.com/stewkk/t-bmstu/pkg/submission"
)

type ProblemDto struct {
	ProblemMetaDto
	Statement string
}

func ToProblemDto(p problem.Problem) ProblemDto {
	return ProblemDto{
		ProblemMetaDto: ToProblemMetaDto(p.ProblemMeta),
		Statement:      p.Statement,
	}
}

type ProblemMetaDto struct {
	Id string
}

func ToProblemMetaDto(p problem.ProblemMeta) ProblemMetaDto {
	return ProblemMetaDto{
		Id: p.Id,
	}
}

type SubmissionMetaDto struct {
	Id string
}

func ToSubmissionMetaDto(s submission.SubmissionMeta) SubmissionMetaDto {
	return SubmissionMetaDto{
		Id: s.Id,
	}
}

type SubmissionInput struct {
	ProblemId string
	Body string
	Language string
}
