package main

import "fmt"

func binarySort(n int, nums []int) []int {
	sorted := []int{}
	index := make(map[int]int)
	og := nums
	for i, n := range nums {
		index[n] = i
	}
	for {
		div := len(nums) / 2
		fmt.Println("div - new:", div)
		if div == 0 || n == nums[div] {
			if n < nums[div] {
				sorted = append(sorted, og[:index[nums[div]]]...)
				sorted = append(sorted, n)
				sorted = append(sorted, og[index[nums[div]]:]...)
			}
			if n >= nums[div] {
				sorted = append(sorted, og[:index[nums[div]]]...)
				sorted = append(sorted, og[index[nums[div]]])
				sorted = append(sorted, n)
				if index[nums[div]] < len(og)-1 {
					sorted = append(sorted, og[index[nums[div]]+1:]...)
				}
			}
			fmt.Println(nums)
			return sorted
		}
		if n < nums[div] {
			fmt.Printf("< nums[%d]\n", div)
			nums = nums[:div]
			fmt.Printf("nums[:%d] = %v\n", div, nums)
			fmt.Println()
			continue
		}
		if n > nums[div] {
			fmt.Printf("> nums[%d]\n", div)
			nums = nums[div:]
			fmt.Printf("nums[%d:] = %v\n", div, nums)
			fmt.Println()
			continue
		}
	}
}

/* func (this *MedianFinder) FindMedian() float64 {
	if len(this.Nums) == 0 {
		return 0
	}
	nums := (this.Nums)
	l := len(nums)
	div := l / 2
	if l%2 == 1 {
		return float64(nums[div])
	}
	lh, rh := nums[:div], nums[div:]
	median := (float64(lh[len(lh)-1]) + float64(rh[0])) / 2.0
	return median
}
*/
