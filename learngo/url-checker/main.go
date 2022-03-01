package main

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrRequestFailed = errors.New("request failed")

type response struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)

	urls := []string{
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.coupang.com/",
		"https://www.airbnb.co.kr/",
		"https://www.goodchoice.kr/",
	}

	urlCheckChannel := make(chan response)

	for _, url := range urls {
		go hitURL(url, urlCheckChannel)
	}

	for i := 0; i < len(urls); i++ {
		resp := <-urlCheckChannel
		results[resp.url] = resp.status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, c chan<- response) { // <- 를 매개변수 타입에 입력하면 채널에서 데이터를 보낼수만 있도록 고정할 수 있다.
	fmt.Println("Checking: ", url)

	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- response{
		url:    url,
		status: status,
	}
}
