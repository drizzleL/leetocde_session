package main

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i { // found
			i++
			continue
		}
		j := nums[i]
		if nums[j] == j {
			return j
		}
		nums[i], nums[j] = nums[j], nums[i] // swap
	}
	return -1
}
