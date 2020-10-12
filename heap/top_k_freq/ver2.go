package main

import (
	"container/heap"
	"fmt"
)

type entry struct {
	val   int
	count int
}

type entries []*entry

func (e entries) Len() int           { return len(e) }
func (e entries) Less(i, j int) bool { return e[i].count > e[j].count }
func (e entries) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func (e *entries) Push(x interface{}) {
	s := x.(entry)
	*e = append(*e, &entry{val: s.val, count: s.count})
}

func (e *entries) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

// sort entries using sort package,
func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]bool)
	em := make(map[int]int)
	var h entries
	heap.Init(&h)
	q := 0
	for i, n := range nums {
		if !seen[n] {
			seen[n] = true
			e := entry{val: n, count: 1}
			heap.Push(&h, e)
			em[n] = i
		} else {
			q = em[n]
			e := entry{val: n, count: h[q].count + 1}
			heap.Push(&h, e)
			fmt.Printf("k: %d v.Val: %d v.Count: %d\n", n, h[q].val, h[q].count)
		}

	}
	fmt.Println(em)

	most := make([]int, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(&h).(*entry).val
		fmt.Println(p)
		most[i] = p
	}
	return most
}

func main() {
	nums := []int{1, 2, 1, 2, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 2))
}
