package main

import "fmt"

/*
	Go 언어는 클래스를 제공하지 않으므로 상속 또한 제공하지 않는다.
	하지만 구조체에서 Embedding을 사용하면 상속과 같은 효과를 낼 수 있다.
*/

type Person2 struct {
	name string
	age  int
}

func (p *Person2) greeting() {
	fmt.Println("Hello world!")
}

type Student struct {
	p      Person2 // 학생 구조체 안에 사람 구조체를 필드로 가지고 있다. (Has-a 관계)
	school string
	grade  int
}
