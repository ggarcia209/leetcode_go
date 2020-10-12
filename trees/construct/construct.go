package construct

// TreeNode represents one node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNode(val int) *TreeNode {
	if val == 0 {
		return &TreeNode{}
	}
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

// CreateTree creates binary search tree from given list
func CreateTree(vals []int) *TreeNode {
	j := 1                           // records # of nodes processed to find subsequent nodes' descendent's indices
	nodes := make(map[int]*TreeNode) // used to link each generation of nodes to the it's parent
	root := &TreeNode{}

	for i, val := range vals {
		// initialize root and its children
		if j == 1 {
			root.Val = val
			nodes[i] = root
			// i = node index, j = distance to first child, left child = i+j, right child = i+j+1
			if i+j < len(vals) {
				root.Left = CreateNode(vals[i+j])
				nodes[i+j] = root.Left
			}
			if i+j+1 < len(vals) {
				root.Right = CreateNode(vals[i+j+1])
				nodes[i+j+1] = root.Right
			}
			j++
			continue
		}

		if i+j > len(vals)-1 { // no more descendents
			return root
		}
		// retreive next node in map and create next nodes children
		node := nodes[i]
		node.Left = CreateNode(vals[i+j])
		nodes[i+j] = node.Left
		if i+j == len(vals)-1 {
			break
		}

		node.Right = CreateNode(vals[i+j+1])
		nodes[i+j+1] = node.Right
		if i+j+1 == len(vals)-1 {
			break
		}

		j++ // increase distance to next node's childrens' indices
	}
	return root
}

// FindNodes returns a list of each node
func FindNodes(root *TreeNode) []int {
	var nodes []int
	if root == nil {
		return nodes
	}

	for {
		nodes
	}
	nodes = append(nodes, root.Val)
	nodes = append(nodes, FindNodes(root.Left)...)
	nodes = append(nodes, FindNodes(root.Right)...)
	return nodes
}
