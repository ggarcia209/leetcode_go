/* Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid. */
package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := createList(5)
	showList(head)
	new := removeNthFromEnd(head, 5)
	showList(new)
}

func createList(n int) *ListNode {
	node := &ListNode{1, nil}
	head := node
	for i := 1; i < n; i++ {
		node.Next = &ListNode{i + 1, nil}
		node = node.Next
	}
	return head
}

func showList(head *ListNode) {
	node := head
	for {
		if node.Next == nil {
			fmt.Printf("node.Val: %d\n", node.Val)
			break
		}
		fmt.Printf("node.Val: %d, node.Next.Val: %d\n", node.Val, node.Next.Val)
		node = node.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make(map[int]*ListNode)
	node := head
	var rm *ListNode
	for i := 1; i > 0; i++ {
		nodes[i] = node
		node = node.Next
		if node == nil {
			if n == i {
				rm = nodes[i-(n-1)]
				rm.Next = nil
				return nodes[i-(n-2)]
			}
			rm = nodes[i-(n)]
			rm.Next = nodes[i-(n-2)]
			break
		}

	}
	return head
}
