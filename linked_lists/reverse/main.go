/* Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL */
package main

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/* func createList(size int) []ListNode {
	var nodes []*ListNode
	for i := 0; i < size; i++ {
		ln := ListNode{i, nil}
		nodes = append(nodes, *ln)
	}
	for i := 0; i < size-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	return nodes
} */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rev := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rev
}
