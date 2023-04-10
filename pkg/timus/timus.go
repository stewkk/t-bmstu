package timus

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stewkk/t-bmstu/pkg/models"
	"net/http"
	"net/url"
)

type TestingSystem struct{}

func (ts *TestingSystem) GetProblem(url string) (models.Problem, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return models.Problem{}, err
	}

	// Find the div with class "problem_content"
	problemContent := doc.Find("div.problem_content")

	// Get the HTML content of the div
	htmlContent, err := problemContent.Html()
	return models.Problem{
		Statement: htmlContent,
	}, err
}

func (ts *TestingSystem) SendSubmission(judge_id string, language string, task_id string, code string) error {
	// This is the function, that send solution to the timus
	url_ := "https://acm.timus.ru/submit.aspx"

	r := url.Values{
		"action":     {"submit"},
		"SpaceID":    {"1"},
		"JudgeID":    {judge_id},
		"Language":   {language},
		"ProblemNum": {task_id},
		"Source":     {code},
	}

	resp, err := http.PostForm(url_, r)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return err
}
