package main

type MinStack struct {
	min  int
	diff []int // val = newmin - oldmin
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.diff) == 0 {
		this.min = x
		this.diff = append(this.diff, 0)
		return
	}
	this.diff = append(this.diff, x-this.min)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	topDiff := this.diff[len(this.diff)-1]
	if topDiff < 0 { // this.min is top
		this.min = this.min - topDiff
	}
	this.diff = this.diff[:len(this.diff)-1]
}

func (this *MinStack) Top() int {
	topDiff := this.diff[len(this.diff)-1]
	if topDiff < 0 {
		return this.min
	}
	return this.min + topDiff
}

func (this *MinStack) Min() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
