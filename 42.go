package main

import "math"

func trap(height []int) int {
	// from left high
	// from right high
	var leftHigh, rightHigh int
	tmp := int(math.MinInt32)
	for i := 0; i < len(height); i++ {
		if height[i] > tmp {
			tmp = height[i]
			leftHigh = i
		}
	}
	tmp = int(math.MinInt32)
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > tmp {
			tmp = height[i]
			rightHigh = i
		}
	}

	var sum int
	for i := 0; i < len(height); i++ {
		sum += height[i]
	}

	var maxSum int
	m := 0
	for i := 0; i < leftHigh; i++ {
		if height[i] > m {
			m = height[i]
		}
		maxSum += m
	}
	m = 0
	for i := len(height) - 1; i > rightHigh; i-- {
		if height[i] > m {
			m = height[i]
		}
		maxSum += m
	}
	maxSum += height[leftHigh] * (rightHigh - leftHigh + 1)
	return maxSum - sum
}
