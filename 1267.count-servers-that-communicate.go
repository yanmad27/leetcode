package main

/*
 * @lc app=leetcode id=1267 lang=golang
 *
 * [1267] Count Servers that Communicate
 */

// @lc code=start
func countServers(grid [][]int) int {

	m, n := len(grid), len(grid[0])
	rows, cols := make([]int, m), make([]int, n)
	computers := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rows[i]++
				cols[j]++
				computers++
			}
		}
	}
	rs := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				rs++
			}
		}
	}
	return computers - rs
}

// @lc code=end
