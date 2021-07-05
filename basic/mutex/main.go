package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 실제로 6개의 물리적 코어를 가지고 있지만
	// 하이퍼스레딩 기술이 사용되어 논리적 코어는 12개이다.
	fmt.Println("CPU 코어 수: ", runtime.NumCPU())

	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	// mutexTest()

	// RWMutex test
	// rwmutexTestX()
	rwmutexTestO()
}
