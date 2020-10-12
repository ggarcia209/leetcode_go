package main

import (
	"fmt"
	"math"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0 // return 0 if list empty
	}

	max := 0                                 // record largest area found
	li, ri := 0, len(height)-1               // index from left, right sides
	w := len(height[:ri]) - len(height[:li]) // initial width/distance between elements

	for li < ri { // iterate until left, right indexes intersect
		w = len(height[:ri]) - len(height[:li])           // new width
		l, r := height[li], height[ri]                    // heights of left, right side
		area := int(math.Min(float64(l), float64(r))) * w // w * min(l,r)

		if l < r {
			li++ // search for larger height from left
		} else {
			ri-- // search for larger height from right
		}
		if area > max {
			max = area
		}
	}
	return max
}
