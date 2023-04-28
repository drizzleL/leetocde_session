package main

import "sort"

func nextPermutation(nums []int) {
	reverseSlice := func(x []int) {
		for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
			x[i], x[j] = x[j], x[i]
		}
	}
	j := len(nums) - 1
	for j != 0 && nums[j-1] >= nums[j] {
		j--
	}
	if j == 0 {
		reverseSlice(nums)
		return
	}
	reverseSlice(nums[j:])
	firstBigger := sort.Search(len(nums[j:]), func(i int) bool {
		return nums[j+i] > nums[j-1]
	})
	nums[j-1], nums[j+firstBigger] = nums[j+firstBigger], nums[j-1]
	return
}
