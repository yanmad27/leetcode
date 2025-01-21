package main

import "math"

/*
 * @lc app=leetcode id=2017 lang=golang
 *
 * [2017] Grid Game
 */

// @lc code=start
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	sum0 := 0
	sum1 := 0
	for i := 0; i < n; i++ {
		sum0 += grid[0][i]
	}
	result := math.MaxInt
	for i := 0; i < n; i++ {
		sum0 -= grid[0][i]
		if i-1 >= 0 {
			sum1 += grid[1][i-1]
		}
		result = min(max(sum0, sum1), result)
	}
	return int64(result)
}

// @lc code=end
