package main

import (
	"container/heap"
	"fmt"
	"time"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *minHeap) View() interface{} {
	n := *h
	return n[0]
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *maxHeap) View() interface{} {
	n := *h
	return n[0]
}

type MedianFinder struct {
	min *minHeap
	max *maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	min, max := &minHeap{}, &maxHeap{}

	return MedianFinder{min, max}
}

func (this *MedianFinder) AddNum(num int) {
	if this.min.Len() < this.max.Len() {
		heap.Push(this.min, num)
		return
	}
	heap.Push(this.max, num)
	return
}

func (this *MedianFinder) FindMedian() float64 {
	for {
		l, r := *this.max, *this.min
		if l[0] > r[0] {
			heap.Init(this.min)
			heap.Init(this.max)
			heap.Pop(this.max)
			heap.Push(this.min, l[0])
			heap.Pop(this.min)
			heap.Push(this.max, r[0])
		} else {
			break
		}
	}
	if this.max.Len() > this.min.Len() {
		return float64(heap.Pop(this.max).(int))
	}
	median := (float64(heap.Pop(this.min).(int) + heap.Pop(this.max).(int))) / 2.0
	return median
}

func main() {
	obj := Constructor()
	obj.AddNum(3)
	obj.AddNum(7)
	obj.AddNum(6)
	obj.AddNum(4)
	obj.AddNum(2)
	obj.AddNum(1)
	obj.AddNum(10)
	// obj.AddNum(5)
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
