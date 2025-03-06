package main

/*
 * @lc app=leetcode id=2965 lang=golang
 *
 * [2965] Find Missing and Repeated Values
 */

// @lc code=start
func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	arr := make([]int, n*n+1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[grid[i][j]]++
		}
	}
	result := []int{0, 0}
	for i := 1; i <= n*n; i++ {
		if arr[i] == 0 {
			result[1] = i 
		} else if arr[i] == 2 {
			result[0] = i 
		}
	}
	return result
}

// @lc code=end
