package main

import (
	"container/heap"
)

/*
 * @lc app=leetcode id=2530 lang=golang
 *
 * [2530] Maximal Score After Applying K Operations
 */

// @lc code=start

// An IntHeap is a min-heap of ints.

func maxKelements(nums []int, k int) int64 {
	h := MaxHeap(nums)
	heap.Init(&h)
	rs := 0
	for i := 1; i <= k; i++ {
		max := heap.Pop(&h)
		rs += max.(int)
		heap.Push(&h, (max.(int)+2)/3)
	}
	return int64(rs)
}

// @lc code=end
