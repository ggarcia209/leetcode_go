package main

import (
	"strings"
)

/*
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters.
The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between
two words. The returned string should only have a single space separating the words.
Do not include any extra spaces.
*/

func main() {
	// ...
}

func reverseWords(s string) string {
	words := strings.Split(strings.Join(strings.Fields(s), " "), " ")
	rev := make([]string, len(words))
	j := len(words) - 1
	for i, w := range words {
		w = strings.TrimSpace(w)
		rev[i] = words[j]
		j--
	}
	res := strings.Join(rev, " ")
	return res
}
