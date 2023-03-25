package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Submission struct {
	id             string
	task_id        string
	language       string
	verdict        string
	number_of_test string
	time           string
	memory         string
}

func parseString(i int, s *goquery.Selection) Submission {
	var current Submission

	// get problem id
	s.Find(".id").Each(func(i int, q *goquery.Selection) {
		current.id = q.Text()
	})

	// get task id
	s.Find(".problem").Each(func(i int, q *goquery.Selection) {
		current.task_id = q.Text()
	})

	// get language
	s.Find(".language").Each(func(i int, q *goquery.Selection) {
		current.language = q.Text()
	})

	flag := 0
	// get verdict
	s.Find(".verdict_rj").Each(func(i int, q *goquery.Selection) {
		current.verdict = q.Text()
		flag = 1
	})

	if flag == 0 {
		current.verdict = "Accepted"
	}

	// get number of test
	s.Find(".test").Each(func(i int, q *goquery.Selection) {
		current.number_of_test = q.Text()
	})

	// get time
	s.Find(".runtime").Each(func(i int, q *goquery.Selection) {
		current.time = q.Text()
	})

	// get memory
	s.Find(".memory").Each(func(i int, q *goquery.Selection) {
		current.memory = q.Text()
	})

	return current
}

func ResultsScrape() []Submission {
	url_ := "https://acm.timus.ru/status.aspx?author=342187&count=100&refresh=0"

	// Request the HTML page.
	res, err := http.Get(url_)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var arr []Submission

	// Find the review items
	doc.Find(".even").Each(func(i int, s *goquery.Selection) {
		arr = append(arr, parseString(i, s))
	})

	doc.Find(".odd").Each(func(i int, s *goquery.Selection) {
		arr = append(arr, parseString(i, s))
	})

	//fmt.Print(arr)
	return arr
}

func SendSubmission(judge_id string, language string, task_id string, code string) {
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
		return
	}

	defer resp.Body.Close()
}

var tasks_map map[uuid.UUID]string

func createMap() {
	tasks_map = make(map[uuid.UUID]string)

	s, _ := uuid.Parse("7a81f316-add0-4dc9-85a7-28b8fd747f8f")
	tasks_map[s] = "https://acm.timus.ru/problem.aspx?space=1&num=1003"

	s, _ = uuid.Parse("0960f297-8e9f-492e-90bd-8708772638e3")
	tasks_map[s] = "https://acm.timus.ru/problem.aspx?space=1&num=1004"

	s, _ = uuid.Parse("e346caa7-5af4-4741-a1a0-d4f5f88a87c2")
	tasks_map[s] = "https://acm.timus.ru/problem.aspx?space=1&num=1005"

	//fmt.Print(tasks_map)
}

func getLink(id uuid.UUID) (string, bool) {
	url_, ok := tasks_map[id]
	return url_, ok
}

func getProblem(id uuid.UUID) {
	url_, ok := getLink(id)

	if !ok {
		return
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
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//createMap()
	// fmt.Println(ResultsScrape())
	//SendSubmission("342187EL", "57", "1000", "a, b = [int(i) for i in input().split()]\nprint(a + b)")
	//a, _ := uuid.Parse("7a81f316-add0-4dc9-85a7-28b8fd747f8f")
	//fmt.Print(getLink(a))
	//getProblem(a)
}
