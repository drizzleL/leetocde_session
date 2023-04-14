package main

import "sort"

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for nums[i]+nums[j] != target {
		delta := sort.Search(j-i+1, func(v int) bool {
			idx := i + v
			return nums[idx]+nums[i] > target
		})
		j = i + delta - 1
		delta = sort.Search(j-i+1, func(v int) bool {
			idx := i + v
			return nums[idx]+nums[j] >= target
		})
		i = i + delta
	}
	return []int{nums[i], nums[j]}
}
