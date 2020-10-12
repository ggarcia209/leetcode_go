package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "A man, a plan, a canal: Panama"
	// str = "race a car"
	// str = "aa"
	is := isPalindrome(str)
	fmt.Println(is)
}

func isPalindrome(s string) bool {
	s = strings.Replace(strings.ToLower(s), " ", "", -1)
	re := regexp.MustCompile("[a-zA-Z]+")
	str := strings.Split(strings.Join(re.FindAllString(s, -1), ""), "")

	buf := []string{}
	for _, s := range str {
		buf = append(buf, strings.ToLower(s))
	}
	fmt.Println(buf)
	j := len(str) - 1
	for _, s := range buf {
		fmt.Println(s, str[j])
		if s != str[j] {
			return false
		}
		j--
	}

	return true
}
