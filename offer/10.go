package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
