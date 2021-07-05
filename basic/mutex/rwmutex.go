package main

import (
	"fmt"
	"sync"
	"time"
)

func rwmutexTestX() {
	var data int = 0

	go func() {
		// 값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			data += 1                         // data에 값 쓰기
			fmt.Println("wirte: ", data)      // data 값 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
		}
	}()

	go func() {
		// 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			fmt.Println("read 1: ", data) // data 값 출력
			time.Sleep(1 * time.Second)   // 1초 대기
		}
	}()

	go func() {
		// 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			fmt.Println("read 2: ", data) // data 값 출력
			time.Sleep(2 * time.Second)   // 2초 대기
		}
	}()

	time.Sleep(10 * time.Second)
}

func rwmutexTestO() {
	var data int = 0
	var rwMutex = new(sync.RWMutex)

	go func() {
		// 값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.Lock() // 쓰기 뮤텍스 잠금

			data += 1                         // data에 값 쓰기
			fmt.Println("wirte: ", data)      // data 값 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기

			rwMutex.Unlock() // 쓰기 뮤텍스 잠금 해제
		}
	}()

	go func() {
		// 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 뮤텍스 잠금

			fmt.Println("read 1: ", data) // data 값 출력
			time.Sleep(1 * time.Second)   // 1초 대기

			rwMutex.RUnlock() // 읽기 뮤텍스 잠금 해제
		}
	}()

	go func() {
		// 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 뮤텍스 잠금

			fmt.Println("read 2: ", data) // data 값 출력
			time.Sleep(2 * time.Second)   // 2초 대기

			rwMutex.RUnlock() // 읽기 뮤텍스 잠금 해제
		}
	}()

	time.Sleep(10 * time.Second)
}
