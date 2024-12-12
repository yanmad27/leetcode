package main

import (
	"container/heap"
	"math"
)

/*
 * @lc app=leetcode id=2558 lang=golang
 *
 * [2558] Take Gifts From the Richest Pile
 */

// @lc code=start

type IntHeap1 []int

func (h IntHeap1) Len() int           { return len(h) }
func (h IntHeap1) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap1) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap1) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) int64 {

	h := IntHeap1(gifts)
	heap.Init(&h)

	for i := 1; i <= k; i++ {
		maxGift := heap.Pop(&h).(int)
		maxGift = int(math.Floor(math.Sqrt(float64(maxGift))))
		heap.Push(&h, int(maxGift))
	}
	sum := 0
	for _, gift := range h {
		sum += gift
	}
	return int64(sum)

}

// @lc code=end
