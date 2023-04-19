package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	ret := []int{}
	for {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			ret = append(ret, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			ret = append(ret, matrix[down][i])
		}
		down--
		if up > down {
			break
		}
		for i := down; i >= up; i-- {
			ret = append(ret, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ret
}
