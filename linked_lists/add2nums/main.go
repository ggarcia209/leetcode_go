package main

import "fmt"


// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	nums1 := []int{1,2,3}
	nums2 := []int{4,5,6,7}
	list1 := nodeGen(nums1)
	list2 := nodeGen(nums2)
	sum := addTwoNumbers(list1, list2)
	printList(sum)
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumRoot *ListNode
	var add *ListNode
	carry := 0
	for {
		next := &ListNode{}
		if add == nil {
			add = &ListNode{}
		}

		// add numbers, carry 1 to next value if current value >= 10
		sum := l1.Val + l2.Val + carry
		add.Val, add.Next = sum, next
		carry = 0  // reset to 0 after carrying 1
		if add.Val > 9 {
			add.Val -= 10
			carry = 1
		}

		if sumRoot == nil {
			sumRoot = add
		}

		// advance to next pair of nodes
		l1, l2 = l1.Next, l2.Next
		if l1 == nil && l2 != nil {  // l2 is longer than l1
			l1 = &ListNode{Val: 0, Next: nil}
		}
		if l2 == nil && l1 != nil {  // l1 is longer than l2
			l2 = &ListNode{Val: 0, Next: nil}
		}
		if l1 == nil && l2 == nil {  // both lists are exhausted
			add.Next = nil
			if carry == 1 {
				add.Next = &ListNode{1, nil}
			}
			// break if at end of both lists
			break
		}
		add = add.Next
	}
	return sumRoot
}

func nodeGen(nums []int) *ListNode {
	nodes := make(map[int]*ListNode)
	var root *ListNode
	for i, num := range nums {
		node := nodes[i]

		if node == nil {
			// create root node; initialize next as empty node
			next := &ListNode{}
			if i+1 > len(nums) - 1 {
				next = nil
			}
			node = &ListNode{Val: num, Next: next}
			root = node
			nodes[i] = node
			nodes[i+1] = next
		}

		next := &ListNode{}
		if i+1 > len(nums) - 1 {
			next = nil
		}
	
		node.Val = num
		node.Next = next
		nodes[i+1] = next
	}
	return root
}

func printList(root *ListNode) {
	i := 0
	curr := root
	for {
		fmt.Printf("Node %d: %v\n", i, curr)
		curr = curr.Next
		if curr == nil {
			break
		}
		i++
	}
}