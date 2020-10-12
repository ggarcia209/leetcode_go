package main

import (
	"fmt"
	"leetcode/graphs/test"
	"leetcode/graphs/util"
)

func main() {
	strIn := test.BigAssInput // true
	// strIn = "[[1,0],[2,0],[3,1],[4,1],[4,2],[5,2]]" // true
	// strIn = "[[1,0],[2,0],[0,3]"                    // true
	// strIn = "[[1,3[,[3,4],[3,1]]"                   // false
	// strIn = "[[1,0],[0,2],[2,1]]"                   // false
	// strIn = "[[]]"                                  // true

	vals := util.ParseInput(strIn)
	can := canFinish(800, vals)

	fmt.Println(can)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := graphGen(prerequisites, numCourses)

	visited := make(map[int]bool)
	curVisited := make(map[int]bool)
	for start := range graph {
		if res := dfs(start, graph, visited, curVisited); !res {
			return false
		}
	}
	return true
}

func dfs(start int, graph map[int][]int, visited map[int]bool, curVisited map[int]bool) bool {
	if val, ok := visited[start]; ok && val {
		return true
	}
	curVisited[start] = true

	neighbors := graph[start]
	for _, neighbor := range neighbors {
		if val, ok := curVisited[neighbor]; ok && val {
			return false
		}
		if res := dfs(neighbor, graph, visited, curVisited); !res {
			return false
		}
	}

	curVisited[start], visited[start] = false, true
	return true
}

// map each classe's prerequisites
func graphGen(prereqs [][]int, numCourses int) map[int][]int {
	graph := make(map[int][]int)
	for _, pair := range prereqs {
		if _, ok := graph[pair[1]]; ok {
			graph[pair[1]] = append(graph[pair[1]], pair[0])
		} else {
			graph[pair[1]] = []int{pair[0]}
		}
	}
	return graph
}
