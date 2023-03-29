package timus

import (
	"github.com/AngelicHedgehog/cf-tool/client"
	"github.com/AngelicHedgehog/cf-tool/config"
	"github.com/AngelicHedgehog/cf-tool/submit"
	"github.com/mitchellh/go-homedir"
	"github.com/stewkk/t-bmstu/pkg/models"
)

type TestingSystem struct{}

func (ts *TestingSystem) GetProblem(problemId models.ProblemId) models.Problem {
	return models.Problem{
		Statement: "TODO",
	}
}

func (ts *TestingSystem) Auth() {
	cfgPath, _ := homedir.Expand("~/.cf/config")
	clnPath, _ := homedir.Expand("~/.cf/session")
	config.Init(cfgPath)
	client.Init(clnPath, config.Instance.Host, config.Instance.Proxy)

	cfg := config.Instance
	cln := client.Instance

	cln.ConfigLogin(false, "serj_1313@mail.ru", "password-haha")
	cfg.AddTemplate(false, "31", "template.py", []string{""}, "py", "", "python3 $%full%$", "", false)
}

func (ts *TestingSystem) Submit() {
	client.Instance.Submit(
		client.Info{
			ProblemType: "contest",
			ContestID:   "896",
			ProblemID:   "B",
		},
		"31", `print('Haha, это работает...')`,
	)
}

func (ts *TestingSystem) GetStatus() string {
	return string(submit.FinalOutput)
}
