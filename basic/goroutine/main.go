package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("say hi")
	fmt.Println(runtime.NumCPU())      // CPU 개수
	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값
	fmt.Println(runtime.GOMAXPROCS(1)) // CPU를 한개만 사용하도록 설정

	// 고루틴은 반복문이 완전히 끝난 다음 생성되므로 고루틴이 생성된 시점의 변수 i의 값은 100이 된다.
	// 따라서 모두 100이 출력된다.
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("say hi3", i)
		}()
	}

	// 클로저를 고루틴으로 실행할때 반복문에 의해 바뀌는 변수는 반드시 매개변수로 넘겨줘야한다.
	// 매개변수로 넘겨주는 시점의 해당 변수의 값이 복사되므로 고루틴이 생성될 때 그대로 사용할 수 있다.
	// for i := 0; i < 100; i++ {
	// 	go func(n int) {
	// 		fmt.Println("say hi2", n)
	// 	}(i)
	// }

	fmt.Scanln()
}