package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 실제로 6개의 물리적 코어를 가지고 있지만
	// 하이퍼스레딩 기술이 사용되어 논리적 코어는 12개이다.
	fmt.Println("CPU 코어 수: ", runtime.NumCPU())

	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			data = append(data, 1)
			mutex.Unlock()

			runtime.Gosched()
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(data))
}
