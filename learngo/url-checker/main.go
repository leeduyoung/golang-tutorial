package main

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrRequestFailed = errors.New("request failed")

func main() {
	var results = make(map[string]string)

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

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)

		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking: ", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return ErrRequestFailed
	}

	return nil
}
