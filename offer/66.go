package main

func constructArr(a []int) []int {
	sum := 1
	ret := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		ret[i] = sum
		sum *= a[i]
	}
	sum = 1
	for i := len(a) - 1; i >= 0; i-- {
		ret[i] *= sum
		sum *= a[i]
	}
	return ret
}
