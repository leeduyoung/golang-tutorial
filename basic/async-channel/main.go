package main

import (
	"fmt"
)

func main() {
	fmt.Println("start async channel")

	/*
		비동기 채널
	*/
	done := make(chan bool, 2)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인 함수 : ", i)
	}

	fmt.Scanln()
}
