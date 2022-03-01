package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=3"

type extractedJob struct {
	id          string
	title       string
	companyName string
	location    string
	summary     string
}

func main() {
	totalCount := getTotalPageCount()
	fmt.Println(totalCount)

	for i := 0; i < totalCount; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	item := doc.Find(".mosaic-provider-jobcards")
	item.Find("a").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("data-jk")
		title, _ := s.Find(".jobTitle>span").Attr("title")
		companyName := s.Find(".companyName").Text()
		companyLocation := s.Find(".companyLocation").Text()
		summary := s.Find(".job-snippet").Text()
		fmt.Println(id, title)
		fmt.Println(companyName, companyLocation, summary)
	})
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
