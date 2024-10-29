package main

/*
 * @lc app=leetcode id=2684 lang=golang
 *
 * [2684] Maximum Number of Moves in a Grid
 */

// @lc code=start
func maxMoves(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	dpAr := make([][]int, m)
	for i := 0; i < m; i++ {
		dpAr[i] = make([]int, n)
		dpAr[i][0] = 1
	}
	maxMove := 1

	for j := 1; j < n; j++ {
		for i := 0; i < m; i++ {
			cur := grid[i][j]
			if i-1 >= 0 && dpAr[i-1][j-1] > 0 && cur > grid[i-1][j-1] {
				dpAr[i][j] = max(dpAr[i-1][j-1]+1, dpAr[i][j])
				maxMove = max(maxMove, dpAr[i][j])
			}
			if i+1 < m && dpAr[i+1][j-1] > 0 && cur > grid[i+1][j-1] {
				dpAr[i][j] = max(dpAr[i+1][j-1]+1, dpAr[i][j])
				maxMove = max(maxMove, dpAr[i][j])
			}
			if dpAr[i][j-1] > 0 && cur > grid[i][j-1] {
				dpAr[i][j] = max(dpAr[i][j-1]+1, dpAr[i][j])
				maxMove = max(maxMove, dpAr[i][j])
			}
		}
	}
	return maxMove - 1

}

// @lc code=end
