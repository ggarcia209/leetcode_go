/* Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3. */

package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 1
	switch {
	case root == nil:
		return 0
	case root.Left == nil && root.Right == nil:
		depth = 1
	}
	dr := maxDepth(root.Right)
	dl := maxDepth(root.Left)
	if dl > dr {
		depth += dl
		return depth
	}
	depth += dr
	return depth
}
