package main

import "fmt"

type Queue struct {
	Elements []int
	Value    int
}

func (q *Queue) EnQueue(val int) []int {
	q.Elements = append(q.Elements, val)
	return q.Elements
}

func (q *Queue) DeQueue(val int) []int {
	q.Elements = q.Elements[1:]
	return q.Elements
}

func main() {
	queue := Queue{Value: 1}
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	fmt.Println(" EnQueue: ", queue)
	queue.DeQueue(5)
	fmt.Println(" DeQueue: ", queue)
}
