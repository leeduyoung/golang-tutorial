package main

import "fmt"

type stack struct {
	data []interface{}
}

func New() *stack {
	return &stack{
		data: []interface{}{},
	}
}

func (s *stack) Push(data interface{}) {
	fmt.Println("push data: ", data)
	s.data = append(s.data, data)
}

func (s *stack) Pop() interface{} {
	lastData := s.data[len(s.data)-1]
	fmt.Println("pop data: ", lastData)

	s.data = s.data[:len(s.data)-1]
	return lastData
}

func (s *stack) get() []interface{} {
	return s.data
}

func main() {
	stack := New()
	stack.Push("10")
	stack.Push("20")
	stack.Push("30")
	stack.Push("40")
	stack.Pop()

	response := stack.get()
	fmt.Println(response...)
}
