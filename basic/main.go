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

	// 구조체의 포인터를 넘김
	rect := Rectangle{30, 30}
	RectangleScaleA(&rect, 10)
	fmt.Println("원래값이 변경됨: ", rect)

	// 구조체의 값을 복사해서 넘김
	rect2 := Rectangle{30, 30}
	RectangleScaleB(rect2, 10)
	fmt.Println("원래값이 변경되지않음: ", rect2)

	rect3 := Rectangle{30, 30}
	rect3.scaleA(10)
	fmt.Println("원래값이 변경됨: ", rect3)

	rect4 := Rectangle{30, 30}
	rect4.scaleB(10)
	fmt.Println("원래값이 변경되지않음: ", rect4)

	// 3. 포인터 변수 사용하기
	PointerTest()

	// 4. 구조체 임베딩을 사용하여 상속하기
	var s Student
	s.p.greeting()
}
