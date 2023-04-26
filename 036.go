package main

func isValidSudoku(board [][]byte) bool {
	getChcker := func() func(i, j int) bool {
		dict := make([]bool, 10)
		return func(i, j int) bool {
			if board[i][j] == '.' {
				return true
			}
			c := board[i][j] - '0'
			if dict[c] {
				return false
			}
			dict[c] = true
			return true
		}
	}
	for i := 0; i < len(board); i++ {
		rowChecker := getChcker()
		for j := 0; j < len(board[0]); j++ {
			if rowChecker(i, j) {
				continue
			}
			return false
		}
	}
	for j := 0; j < len(board[0]); j++ {
		colChecker := getChcker()
		for i := 0; i < len(board); i++ {
			if colChecker(i, j) {
				continue
			}
			return false
		}
	}
	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			regionChecker := getChcker()
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					if regionChecker(i+m, j+n) {
						continue
					}
					return false
				}
			}
		}
	}
	return true
}
