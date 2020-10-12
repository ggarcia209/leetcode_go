package main

import "fmt"

func main() {
	s := ""
	t := "cinemaa"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make(map[string]int)

	for _, l := range s {
		letters[string(l)]++
	}

	for _, l := range t {
		letters[string(l)]--
		if letters[string(l)] == 0 {
			delete(letters, string(l))
		}
	}

	if len(letters) != 0 {
		fmt.Println(letters)
		return false
	}

	return true
}
