package main

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
)

var grid1 = `[["1","0","0","1","1","1","0","1","1","0","0","0","0","0","0","0","0","0","0","0"],["1","0","0","1","1","0","0","1","0","0","0","1","0","1","0","1","0","0","1","0"],["0","0","0","1","1","1","1","0","1","0","1","1","0","0","0","0","1","0","1","0"],["0","0","0","1","1","0","0","1","0","0","0","1","1","1","0","0","1","0","0","1"],["0","0","0","0","0","0","0","1","1","1","0","0","0","0","0","0","0","0","0","0"],["1","0","0","0","0","1","0","1","0","1","1","0","0","0","0","0","0","1","0","1"],["0","0","0","1","0","0","0","1","0","1","0","1","0","1","0","1","0","1","0","1"],["0","0","0","1","0","1","0","0","1","1","0","1","0","1","1","0","1","1","1","0"],["0","0","0","0","1","0","0","1","1","0","0","0","0","1","0","0","0","1","0","1"],["0","0","1","0","0","1","0","0","0","0","0","1","0","0","1","0","0","0","1","0"],["1","0","0","1","0","0","0","0","0","0","0","1","0","0","1","0","1","0","1","0"],["0","1","0","0","0","1","0","1","0","1","1","0","1","1","1","0","1","1","0","0"],["1","1","0","1","0","0","0","0","1","0","0","0","0","0","0","1","0","0","0","1"],["0","1","0","0","1","1","1","0","0","0","1","1","1","1","1","0","1","0","0","0"],["0","0","1","1","1","0","0","0","1","1","0","0","0","1","0","1","0","0","0","0"],["1","0","0","1","0","1","0","0","0","0","1","0","0","0","1","0","1","0","1","1"],["1","0","1","0","0","0","0","0","0","1","0","0","0","1","0","1","0","0","0","0"],["0","1","1","0","0","0","1","1","1","0","1","0","1","0","1","1","1","1","0","0"],["0","1","0","0","0","0","1","1","0","0","1","0","1","0","0","1","0","0","1","1"],["0","0","0","0","0","0","1","1","1","1","0","1","0","0","0","1","1","0","0","0"]]`

type queue struct {
	List [][]int
}

func (q *queue) Check(x, y int, grid *[][]byte) {
	// return when queue exhausted
	if len(q.List) == 0 {
		return
	}

	// check surrounding nodes of current node
	if y > 0 && (*grid)[y-1][x] == 49 { // check above
		q.Add(x, y-1, grid)
	}
	if x < len((*grid)[0])-1 && (*grid)[y][x+1] == 49 { // check right
		q.Add(x+1, y, grid)
	}
	if y < len(*grid)-1 && (*grid)[y+1][x] == 49 { // check below
		q.Add(x, y+1, grid)
	}
	if x > 0 && (*grid)[y][x-1] == 49 { // check left
		q.Add(x-1, y, grid)
	}

	// check surrounding nodes of next node in queue
	x, y = q.Pop() // LIFO
	q.Check(x, y, grid)
}

func (q *queue) Add(x, y int, grid *[][]byte) {
	q.List = append(q.List, []int{x, y}) // LIFO
	(*grid)[y][x] = 48                   // mark as seen/sink island as traversed
}

func (q *queue) Pop() (x, y int) {
	last := q.List[len(q.List)-1]
	x, y = last[0], last[1]
	q.List = q.List[:len(q.List)-1]
	return x, y
}

func numIslands(grid [][]byte) int {
	islands := 0
	var q queue
	for y, row := range grid {
		// fmt.Println(y, row)
		for x, cell := range row {
			// check cells' values
			if cell == 49 {
				islands++
				q.Add(x, y, &grid)
				q.Check(x, y, &grid)
			}
		}
	}
	return islands
}

func main() {
	grid := parseInput(grid1, 20)
	printGrid(grid)
	i := numIslands(grid)
	fmt.Println("islands: ", i)
}

func parseInput(str string, w int) [][]byte {
	grid := [][]byte{}
	var buf bytes.Buffer

	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(str, -1)

	for i, n := range nums {
		io.WriteString(&buf, n)
		if (i+1)%w == 0 {
			row, _ := buf.ReadBytes(2)
			grid = append(grid, row)
			buf.Reset()
		}
	}

	return grid
}

func printGrid(grid [][]byte) {
	x := []int{}

	for i := 0; i < len(grid); i++ {
		x = append(x, i)
	}

	fmt.Printf(" \t%d\n", x)

	for i, row := range grid {
		fmt.Printf("%d\t%d\n", i, row)
	}
}
