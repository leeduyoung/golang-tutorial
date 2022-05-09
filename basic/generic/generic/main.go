package main

import (
	"fmt"
)

type stack[T any] struct {
	data []T
}

type Stack[T any] interface {
	Push(data T)
	Pop() T
	Get() []T
}

func New[T any]() Stack[T] {
	return &stack[T]{
		data: []T{},
	}
}

func (s *stack[T]) Push(data T) {
	fmt.Println("push data: ", data)
	s.data = append(s.data, data)
}

func (s *stack[T]) Pop() T {
	lastData := s.data[len(s.data)-1]
	fmt.Println("pop data: ", lastData)

	s.data = s.data[:len(s.data)-1]
	return lastData
}

func (s *stack[T]) Get() []T {
	return s.data
}

func main() {
	intStack := New[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	intStack.Push(4)
	intStack.Pop()

	response := intStack.Get()
	fmt.Println(response)

	stringStack := New[string]()
	stringStack.Push("a")
	stringStack.Push("b")
	stringStack.Push("c")
	stringStack.Push("d")
	stringStack.Pop()

	response2 := stringStack.Get()
	fmt.Println(response2)
}
