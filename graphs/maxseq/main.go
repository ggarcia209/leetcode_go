package main

import "fmt"

func main() {
	nums := []int{2, 100, 1, 200, 3, 4}
	l := longestConsecutive(nums)
	fmt.Println(l)
}

type node struct {
	val  int
	next int // val + 1
	prev int // val - 1
	seen bool
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	graph := make(map[int]*node)
	for _, n := range nums {
		if graph[n] == nil {
			graph[n] = &node{n, n + 1, n - 1, false}
		}
	}

	curr := 1
	max := curr
	for k := range graph {
		node := graph[k]
		if !node.seen {
			// traverse to beginning of sequence
			for {
				if graph[node.prev] != nil {
					node = graph[node.prev]
				} else if graph[node.prev] == nil {
					break
				}
			}
			// count from beginning of sequence
			for {
				if graph[node.next] != nil {
					node.seen = true // node is only counted in longest possible sequence
					curr++
					node = graph[node.next]
				} else if graph[node.next] == nil {
					break
				}
			}
			if curr > max {
				max = curr
			}

		}
		curr = 1 // reset counter
	}

	return max
}
