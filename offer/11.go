package main

// 查找反转点,二分
func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		mid := l + (r-l)/2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			r--
		}
	}
	return numbers[l]
}
