package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	count := 3

	/*
		동기 채널

		동기 채널은 보내는 쪽에서는 값을 받을때까지 대기하고
		받는쪽에서는 채널에 값이 들어올때까지 대기합니다.

		따라서, 동기 채널을 활용하면 고루틴의 코드 실행 순서를 제어할 수 있다.
	*/
	go func() {
		for i := 0; i < count; i++ {
			done <- true // 고루틴에 true를 보냄. 값을 꺼낼때까지 대기
			fmt.Println("고루틴: ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		<-done // 채널에서 값이 들어올때까지 대기
		fmt.Println("메인 함수: ", i)
	}
}
