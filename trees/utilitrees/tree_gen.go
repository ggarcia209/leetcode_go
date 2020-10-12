// Package utilitrees contains functions for creating binary trees
package utilitrees

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// TreeNode represents one node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Nodes1 is a test input variable
var Nodes1 = []int{1, 2, 3, 4, 5, 0}

// CreateNode creates a binary tree node with the specified value
func CreateNode(val string) *TreeNode {
	if val == "null" {
		return &TreeNode{}
	}
	i, _ := strconv.Atoi(val)
	return &TreeNode{Val: i, Left: nil, Right: nil}
}

// CreateTree creates binary search tree from given list
func CreateTree(vals []string) *TreeNode {
	j := 1
	nodes := make(map[int]*TreeNode)
	root := &TreeNode{}
	for i, val := range vals {
		if j == 1 {
			root.Val, _ = strconv.Atoi(val)
			nodes[i] = root
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
		if i+j > len(vals)-1 {
			return root
		}

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
		j++
	}
	return root
}

// ShowNodes displays info for each node in list
func ShowNodes(nodes []*TreeNode) {
	for i, node := range nodes {
		fmt.Printf("%d.\tVal:%d\tLeft: %v\tRight: %v\n", i, node.Val, node.Left, node.Right)
	}
}

// FindNodes returns a list of each node
func FindNodes(root *TreeNode) []*TreeNode {
	var nodes []*TreeNode
	if root == nil {
		return nodes
	}
	nodes = append(nodes, root)
	nodes = append(nodes, FindNodes(root.Left)...)
	nodes = append(nodes, FindNodes(root.Right)...)
	return nodes
}

// ParseInput parses string input and returns list of node values
func ParseInput(in string) (vals []string) {
	scan := bufio.NewScanner(bytes.NewReader([]byte(in)))
	scan.Split(inSplitFunc)
	for scan.Scan() {
		val := strings.TrimSpace(scan.Text())
		if scan.Text() != "" {
			vals = append(vals, val)
		}
	}
	return vals
}

// inSplitFunc formats input data to be converted to value list
func inSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	str := string(data)
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.Index(str, "["); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if i := strings.Index(str, ","); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if i := strings.Index(str, "]"); i >= 0 {
		return i + 1, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return
}
