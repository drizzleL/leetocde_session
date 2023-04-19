package main

func sumNums(n int) int {
	var res int
	var helper func(x int) int
	helper = func(x int) int {
		_ = x > 1 && helper(x-1) > 0
		res += x
		return res
	}
	return helper(n)
}
