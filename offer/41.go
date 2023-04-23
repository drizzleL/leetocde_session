package main

import "container/heap"

type MedianFinder struct {
	mins MinHeap
	maxs MaxHeap
}

/** initialize your data structure here. */
func ConstructorM() MedianFinder {
	return MedianFinder{
		mins: MinHeap{},
		maxs: MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxs.Len() > this.mins.Len() {
		heap.Push(&this.mins, num)
	} else {
		heap.Push(&this.maxs, num)
	}
	if this.mins.Len() != 0 && this.maxs.Len() != 0 { // balance
		if this.maxs.IntHeap[0] > this.mins.IntHeap[0] { // switch
			bigger := heap.Pop(&this.maxs).(int)
			smaller := heap.Pop(&this.mins).(int)
			heap.Push(&this.maxs, smaller)
			heap.Push(&this.mins, bigger)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxs.Len() > this.mins.Len() {
		return float64(this.maxs.IntHeap[0])
	}
	return float64(this.mins.IntHeap[0]+this.maxs.IntHeap[0]) / 2
}

type MaxHeap struct {
	IntHeap
}

func (v MaxHeap) Less(i, j int) bool {
	return v.IntHeap[i] > v.IntHeap[j]
}

type MinHeap struct {
	IntHeap
}

func (v MinHeap) Less(i, j int) bool {
	return v.IntHeap[i] < v.IntHeap[j]
}
