package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var nSum func(nums []int, target, n int) [][]int
	nSum = func(nums []int, target, n int) [][]int {
		var ret [][]int
		if n == 2 {
			l, r := 0, len(nums)-1
			for l < r {
				sum := nums[l] + nums[r]
				if sum == target {
					ret = append(ret, []int{nums[l], nums[r]})
					r--
					for l < r && nums[r] == nums[r+1] {
						r--
					}
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				} else if sum > target {
					r--
				} else {
					l++
				}
			}
			return ret
		}
		for i := 0; i < len(nums)-1; i++ {
			if i != 0 && nums[i] == nums[i-1] {
				continue
			}
			for _, t := range nSum(nums[i+1:], target-nums[i], n-1) {
				ret = append(ret, append([]int{nums[i]}, t...))
			}
		}
		return ret
	}
	return nSum(nums, target, 4)
}
