package main

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
)

type gridMap map[int]map[int]bool

var grid1 = `[["1","0","0","1","1","1","0","1","1","0","0","0","0","0","0","0","0","0","0","0"],["1","0","0","1","1","0","0","1","0","0","0","1","0","1","0","1","0","0","1","0"],["0","0","0","1","1","1","1","0","1","0","1","1","0","0","0","0","1","0","1","0"],["0","0","0","1","1","0","0","1","0","0","0","1","1","1","0","0","1","0","0","1"],["0","0","0","0","0","0","0","1","1","1","0","0","0","0","0","0","0","0","0","0"],["1","0","0","0","0","1","0","1","0","1","1","0","0","0","0","0","0","1","0","1"],["0","0","0","1","0","0","0","1","0","1","0","1","0","1","0","1","0","1","0","1"],["0","0","0","1","0","1","0","0","1","1","0","1","0","1","1","0","1","1","1","0"],["0","0","0","0","1","0","0","1","1","0","0","0","0","1","0","0","0","1","0","1"],["0","0","1","0","0","1","0","0","0","0","0","1","0","0","1","0","0","0","1","0"],["1","0","0","1","0","0","0","0","0","0","0","1","0","0","1","0","1","0","1","0"],["0","1","0","0","0","1","0","1","0","1","1","0","1","1","1","0","1","1","0","0"],["1","1","0","1","0","0","0","0","1","0","0","0","0","0","0","1","0","0","0","1"],["0","1","0","0","1","1","1","0","0","0","1","1","1","1","1","0","1","0","0","0"],["0","0","1","1","1","0","0","0","1","1","0","0","0","1","0","1","0","0","0","0"],["1","0","0","1","0","1","0","0","0","0","1","0","0","0","1","0","1","0","1","1"],["1","0","1","0","0","0","0","0","0","1","0","0","0","1","0","1","0","0","0","0"],["0","1","1","0","0","0","1","1","1","0","1","0","1","0","1","1","1","1","0","0"],["0","1","0","0","0","0","1","1","0","0","1","0","1","0","0","1","0","0","1","1"],["0","0","0","0","0","0","1","1","1","1","0","1","0","0","0","1","1","0","0","0"]]`

func main() {
	// grid1 := [][]byte{[]byte{1, 1, 1, 1, 0}, []byte{1, 1, 0, 1, 0}, []byte{1, 1, 0, 0, 0}, []byte{0, 0, 0, 0, 0}} // 1, ok
	// grid2 := [][]byte{[]byte{0, 0, 1, 1, 1}, []byte{0, 0, 0, 0, 1}, []byte{0, 0, 0, 0, 1}, []byte{0, 0, 1, 1, 1}} // 1, ok
	// grid3 := [][]byte{[]byte{0, 0, 0, 0, 0}, []byte{0, 0, 1, 1, 0}, []byte{0, 0, 1, 1, 0}, []byte{0, 0, 0, 0, 1}}                        // 2, ok
	// grid4 := [][]byte{[]byte{0, 0, 0, 0, 0}, []byte{0, 0, 0, 0, 0}, []byte{1, 0, 0, 0, 1}, []byte{1, 0, 1, 1, 1}, []byte{1, 1, 1, 1, 1}} // 2, ok
	// fmt.Println(in[0][2])
	// grid := createGrid(in)
	// islands := numIslands(grid4)
	// fmt.Println(islands)
	grid := parseInput(grid1, 20)
	// printGrid(grid)
	i := numIslands(grid)
	fmt.Println("num islands: ", i)
}

/* func numIslands(grid [][]byte) int {
	gridMap = createGrid(grid)
} */

func numIslands(grid [][]byte) int {
	islands := 0
	seen := make(gridMap)
	for y, row := range grid {
		for x, cell := range row {

			if seen[x] == nil {
				seen[x] = make(map[int]bool)
			}
			if seen[x+1] == nil {
				seen[x+1] = make(map[int]bool)
			}

			if cell == 48 {
				continue
			} else { // cell == 1
				switch {

				case !seen[x][y] && (y == 0 || seen[x][y-1] == false): // up not seen or 0
					fmt.Println("new island found")
					fmt.Printf("val: %d\t(%d, %d)\tseen: %v\n", cell, x, y, seen[x][y])
					seen[x][y] = true
					islands++

					if x < len(row)-1 {
						if seen[x+1][y] && (x == 0 || grid[y-1][x-1] == 48) { // connecting sequences of seen nodes
							islands--
							fmt.Println("connected sequences")
							if y < len(grid)-1 && grid[y+1][x] == 49 { // last in row, x+1 == 0
								seen[x][y+1] = true
								fmt.Println("seen (x, y+1) after connecting")
								continue
							}
						}
						if grid[y][x+1] == 49 { // implied && !seen[x+1][y]
							seen[x+1][y] = true
							fmt.Println("seen (x+1, y)")
						}
					}
					if y < len(grid)-1 && grid[y+1][x] == 49 { // last in row, x+1 == 0
						seen[x][y+1] = true
						fmt.Println("seen (x, y+1)")
					}

				default:
					fmt.Println("default")
					fmt.Printf("val: %d\t(%d, %d)\tseen: %v\n", cell, x, y, seen[x][y])
					if x < len(row)-1 {
						if seen[x+1][y] && (x == 0 || grid[y-1][x-1] == 48) { // connecting sequences of seen nodes
							islands--
							fmt.Println("no new island found - connecting sequences")
							if y < len(grid)-1 && grid[y+1][x] == 49 { // last in row, x+1 == 0
								seen[x][y+1] = true
								fmt.Println("seen (x, y+1)")
								continue
							}
						}
						if grid[y][x+1] == 49 {
							seen[x+1][y] = true
							fmt.Println("seen (x+1, y)")
						}
					} else if y < len(grid)-1 && grid[y+1][x] == 49 { // last in row, x+1 == 0
						seen[x][y+1] = true
						fmt.Println("seen (x, y+1)")
					}
				}
			}

		}
	}
	return islands
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
		fmt.Printf("%d\t%s\n", i, row)
	}
}

func testParse() {
	str := "[[0,0,0,1], [0,1,0,1]]"
	grid := parseInput(str, 4)
	fmt.Println("***** testParse *****")
	fmt.Println(grid)
	for _, row := range grid {
		fmt.Println(row)
		for _, b := range row {
			if string(b) == "0" {
				fmt.Println(0)
			}
			if string(b) == "1" {
				fmt.Println(1)
			}
		}
	}
}

/* func check(x, y int, gr grid, dir string) bool {
	m := make(map[string]bool)
	m[dir] = true

	switch {
	case m["start"]:
		if gr[x+1][y] == 1 {

		}
	}
}


func createGrid(grid [][]byte) map[int]map[int]bool {
	xAxis := make(gridMap)
	for y, row := range grid {
		for x, cell := range row {
			if xAxis[x] == nil {
				xAxis[x] = make(map[int]bool)
			}
			fmt.Printf("cell: %d\t(%d, %d)\n", cell, x, y)
			if cell == 0 {
				xAxis[x][y] = false
			} else {
				xAxis[x][y] = true
			}
		}
	}
	return xAxis
}

*/
