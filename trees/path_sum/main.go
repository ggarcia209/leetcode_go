/* Given a binary tree and a sum, determine if the tree has a root-to-leaf path such
that adding up all the values along the path equals the given sum.
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// ...
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil { // edge case - empty tree
		return false
	}
	if root.Left == nil && root.Right == nil {
		if sum != root.Val {
			return false
		}
		return true
	}

	// move down tree to leaves
	// subtract node.Val from var 'sum' for each recursive call
	// var 'sum' == 0 at last recursive call if sum of Vals == target sum
	left := hasPathSum(root.Left, (sum - root.Val))
	right := hasPathSum(root.Right, (sum - root.Val))

	if left || right {
		return true
	}
	return false
}
