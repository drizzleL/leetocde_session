package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var ret int
	minDiff := int(math.MaxInt64)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			diff := int(math.Abs(float64(sum - target)))
			if diff < minDiff {
				minDiff = diff
				ret = sum
			}
			if diff == 0 {
				return sum
			} else if diff > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ret
}
