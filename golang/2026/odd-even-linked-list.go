/* 328. Odd Even Linked List https://leetcode.com/problems/odd-even-linked-list/ Given a singly linked list, group all odd nodes together followed by the even nodes. 
Please note here we are talking about the node number and not the value in the nodes. You should try to do it in place. 
The program should run in O(1) space complexity and O(nodes) time complexity. 
Note: The relative order inside both the even and odd groups should remain as it was in the input. 
The first node is considered odd, the second node even and so on ... 

I/P: 1 → 2 → 3 → 4 → 5
O/P: (1 → 3 → 5) + (2 → 4) */

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return head, nil
	}
	var (
		odd      = head
		evenHead = head.Next
		even     = evenHead
	)
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = nil
	return head, evenHead
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Print(n.Val)
		if n.Next != nil {
			fmt.Print("-> ")
		}
		n = n.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	odd, even := oddEvenList(l1)
	fmt.Print("Odd List:  ")
	printList(odd)

	fmt.Print("Even List: ")
	printList(even)

}
