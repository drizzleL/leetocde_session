package main

import (
	"sort"
)

func search(nums []int, target int) int {
	mid := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if mid == len(nums) || nums[mid] != target {
		return 0
	}
	r := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target+1
	})
	return r - mid
}

func missingNumber(nums []int) int {
	l, r := 0, len(nums)
	var helper func(l, r int) int
	helper = func(l, r int) int {
		if nums[l] != l {
			return l
		}
		if nums[r-1] != r {
			return r
		}
		mid := l + (r-1-l)/2
		if nums[mid] == mid {
			return helper(mid+1, r-1)
		}
		return helper(l+1, mid)
	}
	return helper(l, r)
}
