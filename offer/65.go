package main

func add(a int, b int) int {
	var sum, carry int
	for b != 0 {
		sum = a ^ b
		carry = a & b << 1
		a = sum
		b = carry
	}
	return a
}
