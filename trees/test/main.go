package main

import (
	"leetcode/trees/utilitrees"
)

func main() {
	// in := utilitrees.ParseInput("[1, 2, 3,4, 5, 6, null, 3]")
	// in1 := utiliTrees.ParseInput("[1, 2, 3, null, null, 4, null]")
	in2 := utilitrees.ParseInput("[1]")
	root1 := utilitrees.CreateTree(in2)
	nodes := utilitrees.FindNodes(root1)
	utilitrees.ShowNodes(nodes)
}
