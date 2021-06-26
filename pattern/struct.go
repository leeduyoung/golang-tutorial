package main

type Person struct {
	name string
	age int
}

// 구조체 생성자 패턴
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}