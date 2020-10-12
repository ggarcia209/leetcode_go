package main

import (
	"fmt"
	"leetcode/trees/utilitrees"
)

func main() {
	in := "[5,1,4,null,null,3,6]"              // false
	in = "[3,1,5,null,null,4,6]"               // true
	in = "[10,5,15,null,null,6,20]"            // false
	in = "[3,null,30,10,null,null,15,null,45]" // false
	root := utilitrees.CreateTree(utilitrees.ParseInput(in))
	// fmt.Println(root)
	fmt.Println(isValidBST(root))
}

func isValidBST(root *utilitrees.TreeNode) bool {
	// if empty
	if root == nil {
		return true
	}
	// check validity of left/right child nodes
	if root.Left.Val >= root.Val || root.Val <= root.Left.Val {
		return false
	}
	// traverse down each side of the tree from the root and check validity
	if checkLeft(root.Left, root) && checkRight(root.Right, root) {
		return true
	}
	return false

}

func checkLeft(parent *utilitrees.TreeNode, root *utilitrees.TreeNode) bool {
	// base case
	if parent == nil {
		return true
	}
	// all nodes left of root < root
	if parent.Val >= root.Val {
		fmt.Printf("false: parent (%d) >= root (%d)\n", parent.Val, root.Val)
		return false
	}
	// left child < parent
	if parent.Left != nil && parent.Left.Val >= parent.Val {
		fmt.Printf("false: left (%d) >= parent (%d)\n", parent.Left.Val, parent.Val)
		return false
	}
	// right child > parent
	if parent.Right != nil && parent.Right.Val <= parent.Val {
		fmt.Printf("false: right (%d) <= parent (%d)\n", parent.Right.Val, parent.Val)
		fmt.Println(parent)
		return false
	}
	// traverse next level of tree
	if checkLeft(parent.Left, root) && checkLeft(parent.Right, root) {
		return true
	}

	return false
}

func checkRight(parent *utilitrees.TreeNode, root *utilitrees.TreeNode) bool {
	// base case
	if parent == nil {
		return true
	}
	// all nodes right of root > root
	if parent.Val <= root.Val {
		fmt.Printf("false: parent (%d) <= root (%d)\n", parent.Val, root.Val)
		return false
	}
	// left child < parent
	if parent.Left != nil && parent.Left.Val >= parent.Val {
		fmt.Printf("false: left (%d) >= parent (%d)\n", parent.Left.Val, parent.Val)
		return false
	}
	// right child > parent
	if parent.Right != nil && parent.Right.Val <= parent.Val {
		fmt.Printf("false: right (%d) <= parent (%d)\n", parent.Right.Val, parent.Val)
		fmt.Println(parent)
		return false
	}
	// traverse next level of tree
	if checkRight(parent.Left, root) && checkRight(parent.Right, root) {
		return true
	}

	return false
}
