package main

import "strconv"

func translateNum(num int) int {
	if num == 0 {
		return 1
	}
	var helper func(str string) int
	helper = func(str string) int {
		if len(str) <= 1 {
			return 1
		}
		var cnt int
		if str[0] == '1' || (str[0] == '2' && str[1] >= '0' && str[1] <= '5') {
			cnt = helper(str[2:])
		}
		return helper(str[1:]) + cnt
	}
	str := strconv.Itoa(num)
	return helper(str)
}
