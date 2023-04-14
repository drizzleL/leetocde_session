package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestFunc(t *testing.T) {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := sort.Search(len(x), func(i int) bool {
		return x[i] > 5
	})
	fmt.Println(x[v])
}
