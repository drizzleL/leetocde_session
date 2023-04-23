package main

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	getDigitSum := func(i, j int) int {
		return i/10 + i%10 + j/10 + j%10
	}
	var helper func(i, j int) int
	helper = func(i, j int) int {
		if i >= m || j >= n || vis[i][j] || getDigitSum(i, j) > k {
			return 0
		}
		vis[i][j] = true
		return 1 + helper(i+1, j) + helper(i, j+1)
	}
	return helper(0, 0)
}
