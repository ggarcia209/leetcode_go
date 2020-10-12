package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("str, []str len test:")
	s := "ab c"
	fmt.Printf("s: %v\nlen(s) init: %d\n", s, len(s))
	ss := strings.Split(s, "")
	fmt.Printf("ss: %v\nlen(ss) init: %d\n", ss, len(ss))

	fmt.Println("range loop Test:")
	for _, s := range ss {
		fmt.Println(s)
	}

	fmt.Println("slice test:")
	fmt.Printf("ss[0:]: %s\n", ss[0:])

	fmt.Println("map test:")
	m := map[string]int{
		" ":        len(" "),
		"":         len(""),
		"abcabcbb": len("abcabcbb"),
	}
	for k, v := range m {
		fmt.Printf("k: %s, %d\n", k, v)
	}
	fmt.Printf("m[' '] = %d\n", m[" "])
	fmt.Println("buf test:")
	seen := make(map[string]bool)
	buf := []string{}
	for _, c := range ss {
		if !seen[c] {
			seen[c] = true
			buf = append(buf, c)
		} else {
			// break if char already in buf slice/seen
			break
		}
	}
	fmt.Printf("buf: %v, len(buf): %d", buf, len(buf))
}
