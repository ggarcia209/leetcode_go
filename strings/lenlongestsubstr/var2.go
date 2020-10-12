package main

import (
	"fmt"
	// "leetcode/strings/testin"
	"strings"
	"sync"
)

func main() {
	in := "abccxkcd"
	res := lengthOfLongestSubstring(in)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0 // return 0 for nil input
	}
	ls := make(chan int)  // send lengths of substring from each goroutine to channel
	var wg sync.WaitGroup // wait for all goroutines to finish

	ss := strings.Split(s, "")
	// fmt.Printf("len(ss): %d\n", len(ss))

	for i, _ := range ss {
		wg.Add(1)                  // add goroutine to wait group
		go getLen(ss[i:], ls, &wg) // use loop & slices to determine all possible substrings
	}
	go func() {
		wg.Wait() // wait for all goroutines to finish
		close(ls) // close channel when all goroutines done
	}()
	longest := 0
	// fmt.Printf("%v\n", ls)
	for range ss {
		check := <-ls // receive one message for each char each index in ss slice
		// fmt.Printf("check: %d\nok: %v\n\n", check, ok)
		if check > longest {
			longest = check // longest is greatest value sent thru channel
			// fmt.Printf("longest: %d\n", longest)
		}
	}
	return longest
}

var sema = make(chan struct{}, 40) // concurrency limiting counting semaphore

func getLen(ss []string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()           // signal goroutine finsished to wait group
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // defer token release
	seen := make(map[string]bool)
	buf := []string{}
	for _, c := range ss {
		if !seen[c] { // check if value alread seen in string slice
			seen[c] = true
			buf = append(buf, c) // append each char unique to substring to buffer
			// fmt.Printf("buf: %v, len(buf): %d\n", buf, len(buf))
		} else {
			break // break if char already in buf slice/seen
		}
	}
	// fmt.Printf("sending to channel: %d\n\n", len(buf))
	ch <- len(buf) // send length of substring thru channel
	return
}
