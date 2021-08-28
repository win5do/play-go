package main

import "math"

type MinStack struct {
	data []int
	min  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := this.GetMin()
	if x < min {
		min = x
	}
	this.min = append(this.min, min)
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	l := len(this.data)
	if l == 0 {
		return
	}

	this.data = this.data[:l-1]
	this.min = this.min[:l-1]
}

func (this *MinStack) Top() int {
	return getTop(this.data)
}

func (this *MinStack) GetMin() int {
	return getTop(this.min)
}

func getTop(arr []int) int {
	l := len(arr)
	if l == 0 {
		return math.MaxInt32
	}
	return arr[l-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
