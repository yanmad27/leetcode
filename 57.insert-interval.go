package main

import "slices"

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func canMerged(a, b []int) ([]int, bool) {
	if b[0] < a[0] {
		a, b = b, a
	}
	if a[1] < b[0] {
		return nil, false
	}
	result := []int{0, 0}
	result[0] = min(a[0], b[0])
	result[1] = max(a[1], b[1])
	return result, true
}
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	result := [][]int{}
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]
		if newIn, ok := canMerged(last, curr); ok {
			result = append(result[:len(result)-1], newIn)
		} else {
			result = append(result, curr)
		}
	}
	return result
}

// @lc code=end
