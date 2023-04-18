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

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 2
	for i := 0; i < n-2; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
