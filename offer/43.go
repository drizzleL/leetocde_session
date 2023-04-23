package main

import "math"

func countDigitOne(n int) int {
	if n == 0 {
		return 0
	}
	k := 1
	bit := 0
	for n/k >= 10 {
		k *= 10
		bit++
	}
	a := n / k
	b := n % k
	res := a * helper(bit)
	switch {
	case a > 1:
		res += int(math.Pow10(bit))
	case a == 1:
		res += b + 1
	}
	return res + countDigitOne(b)
}

func helper(k int) int {
	return k * int(math.Pow10(k-1))
}
