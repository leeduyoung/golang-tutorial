package main

import "fmt"

func main() {

	// 1. 생성자 패턴 사용하기
	fmt.Println("main called")

	person := NewPerson("jade", 32)
	fmt.Println(person)

	/*
		2. 함수의 매개변수에 구조체 포인터가 아닌 
		일반적인 형태(구조체 인스턴스)로 넘겨주면
		값이 모두 복사되므로 주의해야한다.
	*/
	rect := Rectangle{30, 30}

	// 구조체의 포인터를 넘김
	RectangleScaleA(&rect, 10) 
	fmt.Println("원래값이 변경됨: ", rect)

	rect2 := Rectangle{30, 30}
	RectangleScaleB(rect2, 10)
	fmt.Println("원래값이 변경되지않음: ", rect2)	
}