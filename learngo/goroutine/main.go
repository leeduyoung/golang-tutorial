package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start main")
	startTime := time.Now()

	// top down programming
	// count("mango")
	// count("kaye")

	// 비동기 처리
	// go count("mango")
	// go count("kaye")

	// 비동기 & 채널 처리
	countChan := make(chan string)
	users := []string{"mango", "kaye"}
	for _, user := range users {
		go count(user, countChan)
	}

	for i := range users {
		result := <-countChan
		fmt.Println(i, result)
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("end main")
	fmt.Printf("실행시간: %s\n", elapsedTime)
}

func count(name string, countChan chan string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("[%s] %d \n", name, i+1)
		time.Sleep(time.Millisecond * 150)
	}

	msg := fmt.Sprintf("%s was done.", name)
	countChan <- msg
}
