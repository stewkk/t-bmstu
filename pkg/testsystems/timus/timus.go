package timus

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type TestingSystem struct{}

type Problem struct {
	Statement string
	TimusId   string
}

func (ts *TestingSystem) GetProblem(url string) (Problem, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return Problem{}, err
	}

	// Find the div with class "problem_content"
	problemContent := doc.Find("div.problem_content")

	// Get the HTML content of the div
	htmlContent, err := problemContent.Html()
	return Problem{
		Statement: htmlContent,
	}, err
}

// SendSubmission This is the function, that sends solution to the Timus judge
func (ts *TestingSystem) SendSubmission(judge_id string, language string, task_id string, code string) error {
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

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	matchFound := false

	// Find rows with the class "status status_nofilter"
	doc.Find("table.status.status_nofilter tr").Each(func(i int, s *goquery.Selection) {
		if matchFound {
			return // Break out of the loop if a match has already been found
		}

		id := strings.TrimSpace(s.Find("td.id").Text())
		coder := strings.TrimSpace(s.Find("td.coder a").Text())
		problem := strings.TrimSpace(s.Find("td.problem a").Text())
		dotIndex := strings.Index(problem, ".")

		if dotIndex != -1 {
			problem = problem[:dotIndex]
		}

		// TODO get name of coder from somewhere
		// TODO maybe do a lang checking?
		if coder == "$tup1d2281337" && problem == task_id {
			matchFound = true
			fmt.Println(id)
		}
	})

	return err
}
