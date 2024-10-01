package main

/*
 * @lc app=leetcode id=1381 lang=golang
 *
 * [1381] Design a Stack With Increment Operation
 */

// @lc code=start
type CustomStack struct {
	stack []int
	size  int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: []int{},
		size:  maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.stack) == this.size {
		return
	}
	this.stack = append(this.stack, x)
}

func (this *CustomStack) Pop() int {
	if len(this.stack) == 0 {
		return -1
	}
	res := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return res
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < len(this.stack) && i < k; i++ {
		this.stack[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
// @lc code=end
