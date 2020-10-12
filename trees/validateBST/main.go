package main

import (
	"fmt"
	"leetcode/trees/utilitrees"
)

func main() {
	in := "[5,1,4,null,null,3,6]"   // false
	in = "[3,1,5,null,null,4,6]"    // true
	in = "[10,5,15,null,null,6,20]" // false
	root := utilitrees.CreateTree(utilitrees.ParseInput(in))
	fmt.Println(root)
	fmt.Println(isValidBST(root))
}

func isValidBST(root *utilitrees.TreeNode) bool {
	if root == nil {
		return true
	}

	switch {
	case root.Left == nil && root.Right == nil:
		return true
	case root.Left == nil:
		if root.Val < root.Right.Val && isValidBST(root.Right) {
			return true
		}
	case root.Right == nil:
		if root.Left.Val < root.Val && isValidBST(root.Left) {
			return true
		}
	case root.Left.Val < root.Val && root.Val < root.Right.Val:
		if root.Left.Right != nil && root.Left.Right.Val > root.Val {
			return false
		}
		if root.Right.Left != nil && root.Right.Left.Val < root.Val {
			return false
		}
		if isValidBST(root.Left) && isValidBST(root.Right) {
			return true
		}
		return false
	default:
		return false
	}
	return false
}
