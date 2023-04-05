package timus

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stewkk/t-bmstu/pkg/models"
	"log"
	"os"
	"strings"
)

type TestingSystem struct{}

var tasks_map map[models.ProblemId]string

func createMap() {
	tasks_map = make(map[models.ProblemId]string)

	//s, _ := uuid.Parse("7a81f316-add0-4dc9-85a7-28b8fd747f8f")
	///tasks_map[s] = "https://acm.timus.ru/problem.aspx?space=1&num=1003"

	///s, _ = uuid.Parse("0960f297-8e9f-492e-90bd-8708772638e3")
	//tasks_map[s] = "https://acm.timus.ru/problem.aspx?space=1&num=1004"

	//s, _ = uuid.Parse("e346caa7-5af4-4741-a1a0-d4f5f88a87c2")
	//tasks_map[s] = "https://acm.timus.ru/problem.aspx?space=1&num=1005"
}

// TODO to a normal getting url
func getLink(id models.ProblemId) (string, bool) {
	url_, ok := tasks_map[id]
	return url_, ok
}

func (ts *TestingSystem) GetProblem(problemId models.ProblemId) models.Problem {
	url_, ok := getLink(problemId)

	if !ok {
		return models.Problem{
			Statement: "error",
		}
	}

	//fmt.Print(url_)
	// get the id of task
	task_id := url_[(strings.Index(url_, "num=") + 4):]

	// русское условие
	url_ += "&locale=ru"
	// Load the HTML document
	doc, err := goquery.NewDocument(url_)
	if err != nil {
		log.Fatal(err)
	}

	// Find the div with class "problem_content"
	problemContent := doc.Find("div.problem_content")

	// Get the HTML content of the div
	htmlContent, err := problemContent.Html()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new file to write the content to
	file, err := os.Create(task_id + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(htmlContent)
	return models.Problem{
		Statement: htmlContent,
	}
}
