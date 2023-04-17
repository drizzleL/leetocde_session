package main

func singleNumbers(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	m := xor & (^xor + 1)
	var x, y int
	for _, num := range nums {
		if num&m != 0 {
			x ^= num
		} else {
			y ^= num
		}
	}
	return []int{x, y}
}
