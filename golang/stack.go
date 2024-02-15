package main

import "fmt"

type Stack struct {
	Elements []int
	Value    int
}

func (s *Stack) Push(val int) []int {
	s.Elements = append(s.Elements, val)
	return s.Elements
}

func (s *Stack) Pop(val int) []int {
	n := len(s.Elements) - 1
	s.Elements = s.Elements[:n]
	return s.Elements
}

func main() {
	queue := Stack{Value: 1}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	fmt.Println(" Push: ", queue)
	queue.Pop(5)
	fmt.Println(" Pop: ", queue)
}
