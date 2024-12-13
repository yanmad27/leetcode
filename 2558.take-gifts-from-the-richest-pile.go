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

func pickGifts(gifts []int, k int) int64 {

	h := MaxHeap(gifts)
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
