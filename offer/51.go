package main

func reversePairs(nums []int) int {
	var merge func(x []int) int
	merge = func(x []int) int {
		if len(x) <= 1 {
			return 0
		}
		mid := len(x) / 2
		left, right := x[:mid], x[mid:]
		res := merge(left) + merge(right)
		newArr := make([]int, 0, len(left)+len(right))
		for len(left) != 0 || len(right) != 0 {
			if len(left) == 0 || len(right) == 0 { // one is empty, push another
				newArr = append(newArr, left...)
				newArr = append(newArr, right...)
				break
			}
			if left[0] <= right[0] {
				newArr = append(newArr, left[0])
				left = left[1:]
				continue
			} else {
				res += len(left)
				newArr = append(newArr, right[0])
				right = right[1:]
				continue
			}
		}
		copy(x, newArr)
		return res
	}
	return merge(nums)
}
