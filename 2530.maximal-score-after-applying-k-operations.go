package main

import (
	"container/heap"
	"math"
)

/*
 * @lc app=leetcode id=2530 lang=golang
 *
 * [2530] Maximal Score After Applying K Operations
 */

// @lc code=start

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

func maxKelements(nums []int, k int) int64 {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}
	rs := 0
	for i := 1; i <= k; i++ {
		max := heap.Pop(h)
		rs += max.(int)
		tmp := float64(max.(int)) / 3
		replace := math.Ceil(tmp)
		heap.Push(h, int(replace))
	}
	return int64(rs)
}

// @lc code=end
