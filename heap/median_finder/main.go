package main

import (
	"fmt"
	"time"
)

type MedianFinder struct {
	Nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.Nums = binarySort(num, this.Nums)
}

func (this *MedianFinder) FindMedian() float64 {

	if len(this.Nums) == 0 {
		return 0
	}
	l := len(this.Nums)
	div := l / 2
	if l%2 == 1 {
		return float64(this.Nums[div])
	}
	median := (float64(this.Nums[div-1]) + float64(this.Nums[div])) / 2.0

	return median
}

func main() {
	obj := Constructor()
	obj.AddNum(3)
	obj.AddNum(7)
	obj.AddNum(6)
	obj.AddNum(4)
	obj.AddNum(6)
	obj.AddNum(1)
	obj.AddNum(2)
	start := time.Now()
	md := obj.FindMedian()
	fmt.Println(time.Since(start))
	fmt.Println(md)
	// nums := []int{2, 4, 6, 8, 10, 12, 14}
	// sorted := binarySort(9, nums)
	// fmt.Printf("sorted = %v\n", sorted)
}

func binarySort(n int, nums []int) []int {
	if len(nums) == 0 {
		return []int{n}
	}
	sorted := []int{}
	index := make(map[int]int)
	og := nums
	for i, n := range nums {
		index[n] = i
	}
	for {
		div := len(nums) / 2
		switch {
		case div == 0 || n == nums[div]:
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
			return sorted
		case n < nums[div]:
			nums = nums[:div]
			continue
		case n > nums[div]:
			nums = nums[div:]
			continue
		}
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
