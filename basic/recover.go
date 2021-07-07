package main

import "fmt"

/*
	Recover

	프로그램이 잘못되어 에러가 발생한 뒤 종료되는 상황을 패닉이라고 한다. 다음과 같이 배열의 크기보다 큰 인덱스에 접근했을때 발생하는 에러가 대표적이다.
	panic 함수는 문법적인 에러는 아니지만 로직에 따라 에러로 처리하고 싶을 때 사용한다.

	recover 함수를 사용하면 패닉이 발생했을때 프로그램이 바로 종료되지 않고 예외처리를 할 수 있으며
	다른 언어의 try/catch 구문과 비슷하게 동작한다.
*/
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
