package main

import "math"

func maxProfit(prices []int) int {
	hold := make([]int, len(prices))
	unhold := make([]int, len(prices))
	unholdCD := make([]int, len(prices))
	hold[0] = -prices[0]
	unholdCD[0] = int(math.MinInt32)
	for i, p := range prices {
		if i != 0 {
			hold[i] = max(unhold[i-1]-p, hold[i-1])
			unhold[i] = max(unholdCD[i-1], unhold[i-1])
			unholdCD[i] = hold[i-1] + p
		}
	}
	return max(unhold[len(prices)-1], unholdCD[len(prices)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
