package sametree

import (
	"fmt"
	"leetcode/trees/utiliTrees"
)

func createNode(val int) *utiliTrees.TreeNode {
	return &utiliTrees.TreeNode{Val: val, Left: nil, Right: nil}
}

func IsSameTree(p *utiliTrees.TreeNode, q *utiliTrees.TreeNode) bool {
	nodesP, nodesQ := FindNodes(p), FindNodes(q)
	if len(nodesP) != len(nodesQ) {
		return false
	}
	for i, node := range nodesP {
		if nodesQ[i] == nil || node == nil {
			if nodesQ[i] != nil || node != nil {
				return false
			}
			continue
		}
		if nodesQ[i].Val != node.Val {
			fmt.Println(nodesQ[i].Val, node.Val)
			return false
		}
		// fmt.Println(node.Val, nodesQ[i].Val)
	}
	return true
}

func FindNodes(root *utiliTrees.TreeNode) (nodes []*utiliTrees.TreeNode) {
	if root == nil {
		nodes = append(nodes, nil)
		return nodes
	}
	nodes = append(nodes, root)
	nodes = append(nodes, FindNodes(root.Left)...)
	nodes = append(nodes, FindNodes(root.Right)...)
	return nodes
}
