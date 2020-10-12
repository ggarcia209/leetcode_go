package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCount := make(map[string]int)
	tCount := make(map[string]int)

	wg.Add(1)
	go count(s, &sCount, &wg)
	wg.Add(1)
	go count(t, &tCount, &wg)

	wg.Wait()

	for _, l := range s {
		if sCount[string(l)] != tCount[string(l)] {
			return false
		}
	}

	return true
}

func count(str string, m *map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, s := range str {
		(*m)[string(s)]++
	}
	return
}

func main() {
	s := "cinem"
	t := "icemn"
	is := isAnagram(s, t)
	fmt.Println(is)
}
