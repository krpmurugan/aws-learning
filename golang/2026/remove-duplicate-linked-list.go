/*
83. Remove Duplicates from Sorted List
https://leetcode.com/problems/remove-duplicates-from-sorted-list/

Given a sorted linked list, delete all duplicates such that each element appear only once.


I/P: 1 → 1 → 2 → 3 → 3
O/p : 1->2->3 

https://github.com/austingebauer/go-leetcode/blob/master/missing_number_268/solution.go
https://github.com/keep-practicing/leetcode-go/blob/master/solutions/0007_reverse_integer/reverse_integer.go
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next // skip duplicate
		} else {
			current = current.Next
		}
	}

	return head
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
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	res := deleteDuplicates(l1)
	printList(res)

}
