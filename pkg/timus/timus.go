package timus

import (
	"github.com/PuerkitoBio/goquery"

	"github.com/stewkk/t-bmstu/pkg/models"
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
