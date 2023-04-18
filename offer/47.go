package main

func maxValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for row := 1; row < m; row++ {
		dp[row][0] = dp[row-1][0] + grid[row][0]
	}
	for col := 1; col < n; col++ {
		dp[0][col] = dp[0][col-1] + grid[0][col]
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			dp[row][col] = grid[row][col] + maxInt(dp[row-1][col], dp[row][col-1])
		}
	}
	return dp[m-1][n-1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
