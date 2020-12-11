package main

/*
There are N gas stations along a circular route, where the amount of gas at
station i is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel
from station i to its next station (i+1). You begin the journey with an empty tank
at one of the gas stations.

Return the starting gas station's index if you can travel around the circuit once
in the clockwise direction, otherwise return -1.

*/

func main() {
	// ...
}

func canCompleteCircuit(gas []int, cost []int) int {
	startPts := []int{}

	for i, n := range gas {
		if n < cost[i] {
			continue
		}
		if i < len(gas)-1 {
			if (n-cost[i])+(gas[i+1]-cost[i+1]) < 0 {
				continue
			}
		} else {
			if (n-cost[i])+(gas[0]-cost[0]) < 0 {
				continue
			}
		}
		startPts = append(startPts, i)
	}

	for _, i := range startPts {
		res := getGas(i, gas, cost)
		if res != -1 {
			return res
		}
	}

	return -1
}

func getGas(start int, gas, cost []int) int {
	tank := 0
	gas = append(gas[start:], gas[:start]...)
	cost = append(cost[start:], cost[:start]...)

	for i, n := range gas {
		tank += n
		tank -= cost[i]
		if tank < 0 {
			return -1
		}
	}

	return start
}
