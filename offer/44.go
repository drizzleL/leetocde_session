package main

import "strconv"

func findNthDigit(n int) int {
	base := 1
	digit := 1
	count := 9
	for n > count {
		n -= count
		base *= 10
		digit += 1
		count = 9 * base * digit
	}
	num := base + (n-1)/digit
	str := strconv.Itoa(num)
	return int(str[(n-1)%digit] - '0')
}
