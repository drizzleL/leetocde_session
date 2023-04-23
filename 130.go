package main

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
			return
		}
		if board[i][j] == '#' || board[i][j] == 'X' {
			return
		}
		board[i][j] = '#'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < len(board); i++ {
		dfs(i, 0)
	}
	for i := 0; i < len(board); i++ {
		dfs(i, len(board[0])-1)
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(0, j)
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(len(board)-1, j)
	}
	// flip boarder
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}
