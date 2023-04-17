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

func findContinuousSequence(target int) [][]int {
	// target = (start+start+delta) * (delta+1) / 2
	var ret [][]int
	for delta := 1; ; delta++ {
		start := ((target * 2 / (delta + 1)) - delta) / 2
		if start == 0 {
			break
		}
		// verify
		if target == (start+start+delta)*(delta+1)/2 {
			tmp := []int{}
			for i := start; i <= start+delta; i++ {
				tmp = append(tmp, i)
			}
			ret = append(ret, tmp)
		}
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}
