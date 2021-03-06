package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type extractedJob struct {
	id              string
	title           string
	companyName     string
	companyLocation string
	summary         string
}

func main() {
	startTime := time.Now()

	var jobs []extractedJob
	totalCount := getTotalPageCount()
	fmt.Println(totalCount)

	for i := 0; i < totalCount; i++ {
		// TODO: 고루틴 처리 가능
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)

	elapsedTime := time.Since(startTime)
	fmt.Println("걸린 시간: ", elapsedTime)
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"id", "title", "companyName", "companyLocation", "summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.companyName, job.companyLocation, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".mosaic-provider-jobcards>a").Each(func(i int, s *goquery.Selection) {
		// TODO: 고루틴 처리 가능
		job := extractJob(s)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(s *goquery.Selection) extractedJob {
	id, _ := s.Attr("data-jk")
	title := cleanString(s.Find(".jobTitle>span").Text())
	companyName := cleanString(s.Find(".companyName").Text())
	companyLocation := cleanString(s.Find(".companyLocation").Text())
	summary := cleanString(s.Find(".job-snippet").Text())
	return extractedJob{
		id:              id,
		title:           title,
		companyName:     companyName,
		companyLocation: companyLocation,
		summary:         summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getTotalPageCount() int {
	count := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		count = s.Find("a").Length()
	})

	return count
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("request failed with status: ", res.StatusCode)
	}
}
