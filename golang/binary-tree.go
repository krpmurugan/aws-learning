package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	preOrder(node.left)
	preOrder(node.right)
}

func main() {
	root := &Node{value: 10}
	root.left = &Node{value: 20}
	root.right = &Node{value: 30}
	root.left.left = &Node{value: 40}
	root.right.right = &Node{value: 50}

	fmt.Println(root)
	preOrder(root)
}
