package main

import (
	"fmt"
)

type Node struct {
	Next  *Node
	value int
}

func NewNode(val int) *Node {
	var n Node
	n.value = val
	n.Next = nil
	return &n
}

func show(head *Node) {
	temp := head
	for temp != nil {
		fmt.Println("->", temp.value)
		temp = temp.Next
	}
}

func main() {
	fmt.Println("Welcome")
	head := NewNode(10)
	head.Next = NewNode(20)
	head.Next.Next = NewNode(30)
	head.Next.Next.Next = NewNode(40)
	show(head)
}
