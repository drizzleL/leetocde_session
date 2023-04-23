package main

import (
	"sort"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j -= 1
		} else {
			i += 1
		}
	}
	return false
}

func find(matrix [][]int, l, r, up, down int) bool {
	if l == r {
		return false
	}
	if up == down {
		return false
	}
	return false
}

// find in arr, >= x
func findX(arr []int, x int) bool {
	idx := sort.SearchInts(arr, x)
	if idx == len(arr) { // 所有都比当前小
		return false
	}
	if arr[idx] == x { // bingo
		return true
	}
	return false
}

func bnSearch(arr []int, x int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] == x {
			return mid
		} else if arr[mid] > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
