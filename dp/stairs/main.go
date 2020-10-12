package main

import "fmt"

func main() {
	steps := climbStairs(2)
	fmt.Println(steps)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1 // only 1 possiblity - 1 step
	}
	dp := make([]int, n+1)    // each slot in dp[1:] represents the total combos for each step
	dp[1], dp[2] = 1, 2       // number of possible combos for first 2 steps
	for i := 3; i <= n; i++ { // start loop at step 3 if n >= 3; iterate thru each step
		// get total possible combos for n steps by adding
		// total number of combos for previous 2 steps (i-1, i-2)
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
