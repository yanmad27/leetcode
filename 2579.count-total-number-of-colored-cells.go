package main

/*
 * @lc app=leetcode id=2579 lang=golang
 *
 * [2579] Count Total Number of Colored Cells
 */

// @lc code=start
func coloredCells(n int) int64 {
	if n == 1 {
		return 1
	}
	return int64(n*n + (n-1)*(n-1))
}

// @lc code=end
