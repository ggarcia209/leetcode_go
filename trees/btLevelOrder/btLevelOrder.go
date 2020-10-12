/* Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7

return its level order traversal as:

[
  [3],
  [9,20],
  [15,7]
] */

package levelorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SuperTreeNode struct {
	*TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// Create queue to iterate thru subsequent child nodes and
	// intialize queue with struct literal representing root of tree
	queue := []SuperTreeNode{SuperTreeNode{TreeNode: root, Level: 0}}
	result := [][]int{}

	for len(queue) != 0 { // while queue is not empty
		node := queue[0]
		level := node.Level
		queue = queue[1:] // pop current node from queue

		if len(result)-1 < level {
			result = append(result, []int{node.Val}) // append level 0 (root) value to result list
		} else {
			result[level] = append(result[level], node.Val) // append current node value to current level list
		}

		// Push child nodes to queue and increment level
		if node.Left != nil {
			queue = append(queue, SuperTreeNode{TreeNode: node.Left, Level: level + 1})
		}
		if node.Right != nil {
			queue = append(queue, SuperTreeNode{TreeNode: node.Right, Level: level + 1})
		}
	}
	return result
}
