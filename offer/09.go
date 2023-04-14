package main

type CQueue struct {
	in, out []int
}

// func Constructor() CQueue {
// 	return CQueue{}
// }

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) > 0 {
		v := this.out[len(this.out)-1]
		this.out = this.out[:len(this.out)-1]
		return v
	}
	if len(this.in) == 0 {
		return -1
	}
	for len(this.in) > 0 {
		v := this.in[len(this.in)-1]
		this.out = append(this.out, v)
		this.in = this.in[:len(this.in)-1]
	}
	return this.DeleteHead()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
