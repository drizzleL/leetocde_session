package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ret, curr int
	for _, num := range nums {
		v1, v2 := num, curr+num
		curr = max(v1, v2)
		ret = max(ret, curr)
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
