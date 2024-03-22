package main

import "fmt"

type Node struct {
	next  *Node
	value int
}

func NewNode(val int) *Node {
	var node Node
	node.next = nil
	node.value = val
	return &node
}

func Travese(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf(" Linked List %d \n", temp.value)
		temp = temp.next
	}

}

func deleteDuplicates(head *Node) *Node {
	current := head

	for current != nil {
		// Check if the next node exists and has the same value
		for current.next != nil && current.value == current.next.value {
			current.next = current.next.next
		}
		current = current.next
	}

	return head
}

func Remove(head *Node) {
	if head == nil {
		return
	}
	temp := head

	for temp.next != nil {
		temp = temp.next
		fmt.Printf(" RLinked List %d \n", temp.value)
	}
	Remove(temp.next)
}

func main() {
	node := NewNode(10)
	node.next = NewNode(20)
	node.next.next = NewNode(20)
	node.next.next.next = NewNode(30)
	node.next.next.next.next = NewNode(40)
	node.next.next.next.next.next = NewNode(40)
	//Travese(node)
	//Remove(node)

	deleteDuplicates(node)
	Travese(node)
	//fmt.Println(head)
}
