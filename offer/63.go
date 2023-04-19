package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	holdMin := prices[0]
	var ret int
	for i := 1; i < len(prices); i++ {
		p := prices[i]
		if p-holdMin > ret {
			ret = p - holdMin
		}
		if p < holdMin {
			holdMin = p
		}
	}
	return ret
}
