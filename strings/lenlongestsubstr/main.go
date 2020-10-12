package main

import (
	"fmt"
	"strings"
)

/* Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
			 Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

TestCase 4: "abccxkcd"
*/

func main() {
	in := "abccxkcd"
	res := lengthOfLongestSubstring(in)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	// return 0 for nil input
	if s == "" {
		return 0
	}
	ss := strings.Split(s, "")
	seen := make(map[string]bool)
	buf := []string{}

	// iterate thru string slice and add chars to buf slice if not seen
	for _, c := range ss {
		if !seen[c] {
			seen[c] = true
			buf = append(buf, c)
		} else {
			// break if char already in buf slice/seen
			break
		}
	}

	this := len(buf) // result from *this* call
	fmt.Printf("buf: %s\nlen(buf): %d\n\n", buf, len(buf))
	// search for longer slice recursively if longer substring possibly exists
	if len(ss[this:]) > this {
		s = strings.Join(ss[this:], "")
		recur := lengthOfLongestSubstring(s) // result from subsequent recursive calls
		if recur > this {
			return recur
		}
	}
	return this
}
