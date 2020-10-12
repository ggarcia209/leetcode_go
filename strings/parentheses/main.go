package main

import "fmt"

var chars = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func main() {
	is := isValid("{}[]")
	fmt.Println(is)
}

func isValid(s string) bool {
	if len(s)%2 != 0 { // if odd number of items
		return false
	}

	seq := []string{} // LIFO queue of right-handed characters
	for _, c := range s {
		if chars[string(c)] != "" { // if left-handed character
			seq = append(seq, chars[string(c)]) // add opposite to queue
		} else { // if right-handed
			if len(seq) == 0 || string(c) != seq[len(seq)-1] { // check current item != most recent item in queue
				return false
			}
			seq = seq[:len(seq)-1] // pop most recent item
		}
	}
	return true
}
