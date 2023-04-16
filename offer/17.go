package main

import "math"

func printNumbers(n int) []int {
	if n == 0 {
		return nil
	}
	size := math.Pow10(n)
	ret := make([]int, int(size))
	for i := 1; i < len(ret); i++ {
		ret[i] = ret[i-1] + 1
	}
	return ret[1:]
}
