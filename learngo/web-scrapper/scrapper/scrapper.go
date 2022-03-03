package scrapper

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

type extractedJob struct {
	id              string
	title           string
	companyName     string
	companyLocation string
	summary         string
}

func Scrape(term string) {
	startTime := time.Now()

	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobChannel = make(chan []extractedJob)
	var jobs []extractedJob
	totalCount := getTotalPageCount(baseURL)
	fmt.Println(totalCount)

	for i := 0; i < totalCount; i++ {
		go getPage(i, jobChannel, baseURL)
	}

	for i := 0; i < totalCount; i++ {
		jobs = append(jobs, <-jobChannel...)
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

func getPage(page int, c chan []extractedJob, baseURL string) {
	var jobChannel = make(chan extractedJob)
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".mosaic-provider-jobcards>a")
	searchCards.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, jobChannel)
	})

	for i := 0; i < searchCards.Length(); i++ {
		jobs = append(jobs, <-jobChannel)
	}

	c <- jobs
}

func extractJob(s *goquery.Selection, c chan<- extractedJob) {
	id, _ := s.Attr("data-jk")
	title := CleanString(s.Find(".jobTitle>span").Text())
	companyName := CleanString(s.Find(".companyName").Text())
	companyLocation := CleanString(s.Find(".companyLocation").Text())
	summary := CleanString(s.Find(".job-snippet").Text())

	c <- extractedJob{
		id:              id,
		title:           title,
		companyName:     companyName,
		companyLocation: companyLocation,
		summary:         summary,
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getTotalPageCount(baseURL string) int {
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
