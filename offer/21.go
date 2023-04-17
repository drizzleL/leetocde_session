package main

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]%2 == 1 {
			l++
			continue
		}
		nums[l], nums[r] = nums[r], nums[l]
		r--
	}
	return nums
}
