package main

type MaxQueue struct {
	vals []int
	maxs []int
}

func ConstructorQ() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.vals) == 0 {
		return -1
	}
	return this.maxs[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.vals = append(this.vals, value)
	for len(this.maxs) != 0 && this.maxs[len(this.maxs)-1] > value {
		this.maxs = this.maxs[:len(this.maxs)-1]
	}
	this.maxs = append(this.maxs, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.vals) == 0 {
		return -1
	}
	val := this.vals[0]
	this.vals = this.vals[1:]
	if val == this.maxs[0] {
		this.maxs = this.maxs[1:]
	}
	return val
}
