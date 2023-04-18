package main

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	x, y, z := 1, 1, 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		num2, num3, num5 := dp[x]*2, dp[y]*3, dp[z]*5
		min := minInt(num2, num3, num5)
		dp[i] = min
		if min == num2 {
			x++
		}
		if min == num3 {
			y++
		}
		if min == num5 {
			z++
		}
	}
	return dp[n]
}

func minInt(x int, nums ...int) int {
	for _, num := range nums {
		if num < x {
			x = num
		}
	}
	return x
}
