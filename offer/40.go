package main

import (
	"container/heap"
)

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := IntHeap(arr[:k])
	heap.Init(&h)

	for i := k; i < len(arr); i++ {
		v := arr[i]
		if v >= h[0] {
			continue
		}
		heap.Pop(&h)
		heap.Push(&h, v)
	}
	return h
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
