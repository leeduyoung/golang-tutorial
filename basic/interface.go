package main

import "fmt"

/*
	인터페이스는 메서드 집합이다. 인터페이스는 메서드 자체를 구현하진 않는다.

	아래 코드는 Printer 인터페이스가 Print라는 메서드를 정의하고 있다.
	MyInt형 변수와 Rectangle형 변수는 Print 메서드를 구현하고 있다.

	인터페이스는 자료형이든 구조체든 타입에 상관없이 메서드 집합만 같으면 동일한 타입으로 보기 때문에
	인터페이스에 대입할 수 있다.
*/

type MyInt int // int형을 MyInt로 정의

// MyInt에 Print 메서드를 연결
func (i MyInt) Print() {
	fmt.Println(i)
}

// 사각형 구조체 정의
type Rectangle2 struct {
	width, height int
}

// Rectangle에 Print 메서드를 연결
func (r Rectangle2) Print() {
	fmt.Println(r.width, r.height)
}

// Print 메서드를 가지는 인터페이스 정의
type Printer interface {
	Print()
}
