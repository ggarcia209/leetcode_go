/* Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	l2 := &ListNode{0, &ListNode{2, &ListNode{5, nil}}}
	start := time.Now()
	root := mergeTwoLists(l1, l2)
	fmt.Println(time.Since(start))
	fmt.Println(root.Val, root.Next.Val)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var node *ListNode
	var root *ListNode
	// determine root
	switch {
	case l1 == nil && l2 == nil:
		return nil
	case l1 == nil:
		node = &ListNode{l2.Val, nil}
		root = node
		l2 = l2.Next
	case l2 == nil:
		node = &ListNode{l1.Val, nil}
		root = node
		l1 = l1.Next
	case l1.Val > l2.Val:
		node = &ListNode{l2.Val, nil}
		root = node
		l2 = l2.Next
	default:
		node = &ListNode{l1.Val, nil}
		root = node
		l1 = l1.Next
	}

L:
	for {
		switch {
		case l1 == nil && l2 == nil: // no remaining nodes in both lists
			break L
		case l1 == nil || l2 == nil: // no remaining nodes in shorter list
			// append remaining nodes of longer list
			// (no further modifications to Next pointers)
			if l1 != nil {
				node.Next = &ListNode{l1.Val, l1.Next}
				break L
			} else {
				node.Next = &ListNode{l2.Val, l2.Next}
				break L
			}
		// determine next node and advance
		case l1.Val > l2.Val:
			node.Next = &ListNode{l2.Val, nil}
			node = node.Next
			l2 = l2.Next
			continue L
		default:
			node.Next = &ListNode{l1.Val, nil}
			node = node.Next
			l1 = l1.Next
			continue L
		}
	}
	return root
}
