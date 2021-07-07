package main

import (
	"fmt"
	"log"
	"time"
)

type HelloOneError struct {
	time  time.Time // 시간
	value int       // 에러를 발생시킨 값
}

// 에러 메시지를 생성하여 리턴하는 Error 함수 구현
func (e HelloOneError) Error() string {
	return fmt.Sprintf("%v: %d는 1이 아닙니다.", e.time, e.value)
}

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	// 에러 타입이 아니더라도 Error 함수를 구현하면 에러로 사용할 수 있다.
	return "", HelloOneError{time.Now(), n}
}

func main() {
	fmt.Println("Hello, world!")

	s, err := helloOne(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	s, err = helloOne(2)
	if err != nil {
		// err로 리턴된 HelloOneError 구조체는 다음과 같이 Type assertion을 해주면 구조체 필드에 접근할 수 있다.
		log.Fatal(err, err.(HelloOneError).value)
	}
	fmt.Println(s)

	fmt.Println("End")
}
