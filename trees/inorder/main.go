package main

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// ...
}

func inorderTraversal(root *TreeNode) []int {
	// left -> root -> right
	nums := []int{}
	if root == nil { // edge case
		return nums
	}

	if root.Left != nil {
		nums = append(nums, inorderTraversal(root.Left)...)
	}

	nums = append(nums, root.Val)

	if root.Right != nil {
		nums = append(nums, inorderTraversal(root.Right)...)
	}

	return nums
}
