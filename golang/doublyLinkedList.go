package main

import "fmt"

type Node struct {
	prev  *Node
	value int
	next  *Node
}

func createNode(value int) *Node {
	var node Node
	node.next = nil
	node.value = value
	node.prev = nil
	return &node
}

func Show(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("=>%d ", temp.value)
		temp = temp.next
	}
}

func main() {
	head := createNode(10)
	node_2 := createNode(20)
	node_3 := createNode(30)
	node_4 := createNode(40)
	head.next = node_2
	node_2.prev = head
	node_2.next = node_3
	node_3.prev = node_2
	node_3.next = node_4
	node_4.prev = node_3
	Show(head)

}
