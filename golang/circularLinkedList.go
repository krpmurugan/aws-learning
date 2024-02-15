package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func show(head *Node) {
	current := head
	elem := current.value
	for {
		fmt.Printf("%d => ", current.value)
		current = current.next
		if current == head {
			break
		}
	}
	fmt.Printf("%d  ", elem)
}

func main() {
	head := &Node{value: 10}
	head.next = &Node{value: 20}
	head.next.next = &Node{value: 30}
	head.next.next.next = head
	show(head)
}
