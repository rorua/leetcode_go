package min_stack

import "math"

type MinStack struct {
	Len  int
	Mins []int
	Vals []int
}

func Constructor() MinStack {
	st := new(MinStack)
	st.Vals = make([]int, 0)
	st.Mins = make([]int, 0)
	return *st
}

func (this *MinStack) Push(val int) {
	this.Vals = append(this.Vals, val)
	if val < this.GetMin() {
		this.Mins = append(this.Mins, val)
	} else {
		this.Mins = append(this.Mins, this.GetMin())
	}
	this.Len++
	return
}

func (this *MinStack) Pop() {
	this.Vals = this.Vals[:this.Len-1]
	this.Mins = this.Mins[:this.Len-1]
	this.Len--
	return
}

func (this *MinStack) Top() int {
	return this.Vals[this.Len-1]
}

func (this *MinStack) GetMin() int {
	if this.Len == 0 {
		return math.MaxInt64
	}
	return this.Mins[this.Len-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
