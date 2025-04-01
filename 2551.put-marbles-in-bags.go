package main

import "sort"

/*
 * @lc app=leetcode id=2551 lang=golang
 *
 * [2551] Put Marbles in Bags
 */

// @lc code=start
func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	adjustmentSum := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		adjustmentSum[i] = weights[i] + weights[i+1]
	}
	sort.Ints(adjustmentSum)
	maxSum, minSum := 0, 0
	for i := range k - 1 {
		maxSum += adjustmentSum[n-2-i]
		minSum += adjustmentSum[i]
	}

	return int64(maxSum - minSum)
}

// @lc code=end
