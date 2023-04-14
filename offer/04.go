package main

import (
	"log"
	"sort"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {
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

func main() {
	v := sort.SearchInts([]int{1, 2, 8, 9}, 15)
	log.Println(v)
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
