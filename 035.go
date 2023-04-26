package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for {
		if target <= nums[l] {
			return l
		}
		if target > nums[r] {
			return r + 1
		}
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
}
