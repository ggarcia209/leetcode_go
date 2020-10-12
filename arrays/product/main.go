package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 4}
	out := productExceptSelf(nums)
	fmt.Println(out)
}

func productExceptSelf(nums []int) []int {
	max := 1
	output := make([]int, len(nums))
	zero := false

	for _, n := range nums {
		if n == 0 && !zero {
			zero = true
			continue
		}
		max *= n  // if more than one 0 in list, max = 0; else max = 1 *= n
	}

	for i, n := range nums {
		if zero && n != 0 {
			output[i] = 0
		} else if zero && n == 0 {
			output[i] = max
		} else {
			output[i] = int(float64(max) * (math.Pow(float64(n), float64(-1))))
		}
	}

	return output
}
