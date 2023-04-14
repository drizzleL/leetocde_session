package main

import (
	"math"
)

func isStraight(nums []int) bool {
	dict := [14]bool{}
	lowest, highest := math.MaxInt, 0
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if dict[num] {
			return false
		}
		dict[num] = true
		if num < lowest {
			lowest = num
		}
		if num > highest {
			highest = num
		}
	}
	if highest == 0 {
		return true
	}
	return highest-lowest <= 4
}
