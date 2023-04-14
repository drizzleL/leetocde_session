package main

func majorityElement(nums []int) int {
	var votes, res int
	for _, num := range nums {
		if votes == 0 {
			votes++
			res = num
		} else if num == res {
			votes++
		} else {
			votes--
		}
	}
	return res
}
