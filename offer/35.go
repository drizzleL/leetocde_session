package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
