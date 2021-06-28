package main

import "fmt"

func recoverTest() {
	defer func() {
		// recover 함수로 런타임 에러 상황을 복구
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2} // 정수 2개가 저장된 배열

	// 배열을 벗어난 접근을 하여 런타임 에러가 발생했지만 recover 함수가 패닉 상황을 복구
	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}
