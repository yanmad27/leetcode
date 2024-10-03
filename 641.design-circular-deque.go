package main

// /*
//  * @lc app=leetcode id=641 lang=golang
//  *
//  * [641] Design Circular Deque
//  */

// // @lc code=start
// type MyCircularDeque struct {
// 	size  int
// 	queue []int
// }

// func Constructor(k int) MyCircularDeque {
// 	return MyCircularDeque{
// 		size:  k,
// 		queue: []int{},
// 	}
// }

// func (this *MyCircularDeque) InsertFront(value int) bool {
// 	if this.size == len(this.queue) {
// 		return false
// 	}
// 	this.queue = append([]int{value}, this.queue...)
// 	return true
// }

// func (this *MyCircularDeque) InsertLast(value int) bool {
// 	if this.size == len(this.queue) {
// 		return false
// 	}
// 	this.queue = append(this.queue, value)
// 	return true
// }

// func (this *MyCircularDeque) DeleteFront() bool {
// 	if len(this.queue) == 0 {
// 		return false
// 	}
// 	this.queue = this.queue[1:]
// 	return true
// }

// func (this *MyCircularDeque) DeleteLast() bool {
// 	if len(this.queue) == 0 {
// 		return false
// 	}
// 	this.queue = this.queue[:len(this.queue)-1]
// 	return true
// }

// func (this *MyCircularDeque) GetFront() int {
// 	if len(this.queue) == 0 {
// 		return -1
// 	}
// 	return this.queue[0]
// }

// func (this *MyCircularDeque) GetRear() int {
// 	length := len(this.queue)
// 	if length == 0 {
// 		return -1
// 	}
// 	return this.queue[len(this.queue)-1]
// }

// func (this *MyCircularDeque) IsEmpty() bool {
// 	return len(this.queue) == 0

// }

// func (this *MyCircularDeque) IsFull() bool {
// 	return len(this.queue) == this.size
// }

// /**
//  * Your MyCircularDeque object will be instantiated and called as such:
//  * obj := Constructor(k);
//  * param_1 := obj.InsertFront(value);
//  * param_2 := obj.InsertLast(value);
//  * param_3 := obj.DeleteFront();
//  * param_4 := obj.DeleteLast();
//  * param_5 := obj.GetFront();
//  * param_6 := obj.GetRear();
//  * param_7 := obj.IsEmpty();
//  * param_8 := obj.IsFull();
//  */
// // @lc code=end
