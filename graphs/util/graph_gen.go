package util

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// ParseInput parses string input and returns list of node values
func ParseInput(in string) [][]int {
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(in, -1)
	pair := []int{}
	vals := [][]int{}

	for _, numStr := range nums {
		val, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			log.Fatal(err)
		}

		pair = append(pair, val)

		if len(pair) == 2 {
			vals = append(vals, pair)
			pair = []int{}
		}
	}

	return vals
}
