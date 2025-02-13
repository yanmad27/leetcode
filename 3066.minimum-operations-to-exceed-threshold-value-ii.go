package main

import "container/heap"

/*
 * @lc app=leetcode id=3066 lang=golang
 *
 * [3066] Minimum Operations to Exceed Threshold Value II
 */

// @lc code=start
func minOperations(nums []int, k int) int {
	h := MinHeap(nums)
	heap.Init(&h)
	num1 := heap.Pop(&h).(int)
	count := 0
	for num1 < k {
		count++
		num2 := heap.Pop(&h).(int)
		newNum := num1*2 + num2
		heap.Push(&h, newNum)
		num1 = heap.Pop(&h).(int)
	}
	return count
}

// @lc code=end
