package main

import (
	"fmt"
	"net/http"

	"github.com/AngelicHedgehog/cf-tool/client"
	"github.com/AngelicHedgehog/cf-tool/config"
	"github.com/AngelicHedgehog/cf-tool/submit"
	"github.com/PuerkitoBio/goquery"
	"github.com/mitchellh/go-homedir"
	"github.com/stewkk/t-bmstu/pkg/models"
)

type TestingSystem struct{}

func (ts *TestingSystem) GetProblem(url string) (models.Problem, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Problem{}, err
	}

	req.Header.Add("Accept-Language", "ru-RU")

	resp, err := client.Do(req)
	if err != nil {
		return models.Problem{}, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return models.Problem{}, err
	}

	// Find the div with class "problem_content"
	problemContent := doc.Find("div.problem-statement")

	// Get the HTML content of the div
	htmlContent, err := problemContent.Html()

	return models.Problem{
		Statement: htmlContent,
	}, err
}

func main() {
	var test TestingSystem

	fmt.Println(test.GetProblem("https://codeforces.com/problemset/problem/1820/B"))
}

func (ts *TestingSystem) Auth(login, password string) {
	cfgPath, _ := homedir.Expand("/config")
	clnPath, _ := homedir.Expand("/session")
	config.Init(cfgPath)
	client.Init(clnPath, config.Instance.Host, config.Instance.Proxy)

	// config.Instance.ConfigLogin(false, login, password)
	// client.Instance.AddTemplate(false, "31", "template.py", []string{""}, "py", "", "python3 $%full%$", "", false)
}

func (ts *TestingSystem) Submit(ContestID, ProblemID, source_code string) {
	client.Instance.Submit(
		client.Info{
			ProblemType: "contest",
			ContestID:   ContestID,
			ProblemID:   ProblemID,
		},
		"31", source_code,
	)
}

func (ts *TestingSystem) GetStatus() string {
	return string(submit.FinalOutput)
}
