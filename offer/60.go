package main

import (
	"math"
)

func dicesProbability(n int) []float64 {
	if n == 0 {
		return nil
	}
	totalSize := 6 * n
	b := make([]float64, 6, totalSize)
	for i := 0; i < 6; i++ {
		b[i] = 1
	}
	sum := func(vv []float64) float64 {
		var ret float64
		for _, v := range vv {
			ret += v
		}
		return ret
	}
	for i := 1; i < n; i++ {
		b = append(b, 0, 0, 0, 0, 0, 0) // add six zero
		for j := len(b) - 1; j >= 0; j-- {
			// 前6个和加到上面
			start := j - 6
			if start < 0 {
				start = 0
			}
			b[j] = sum(b[start:j])
		}
	}

	dividend := math.Pow(6, float64(n))
	for i := range b {
		b[i] = b[i] / float64(dividend)
	}
	return b[n-1:]
}
