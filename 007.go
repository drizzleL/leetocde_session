package main

import "math"

func reverse(x int) int {
	if x < 0 {
		return -1 * reverse(-x)
	}
	ret := 0
	for x != 0 {
		v := x % 10
		x /= 10
		ret *= 10
		ret += v
	}
	if ret > int(math.MaxInt32) {
		return 0
	}
	return ret
}
