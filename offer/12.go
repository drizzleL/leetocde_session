package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	vis := make([][]bool, len(board))
	for i := range board {
		vis[i] = make([]bool, len(board[0]))
	}
	dirs := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	var helper func(i, j int, x int) bool
	helper = func(i, j int, x int) bool {
		if x == len(word) {
			return true
		}
		if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
			return false
		}
		if board[i][j] != word[x] {
			return false
		}
		if vis[i][j] { // seen before
			return false
		}
		vis[i][j] = true
		for _, dir := range dirs {
			if helper(i+dir[0], j+dir[1], x+1) {
				return true
			}
		}
		vis[i][j] = false
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if helper(i, j, 0) {
				return true
			}
		}
	}
	return false
}
