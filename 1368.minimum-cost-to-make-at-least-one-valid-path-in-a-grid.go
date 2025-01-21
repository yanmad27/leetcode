package main

import "math"

/*
 * @lc app=leetcode id=1368 lang=golang
 *
 * [1368] Minimum Cost to Make at Least One Valid Path in a Grid
 */

// @lc code=start
func minCost(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	moves := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	costGrid := make([][]int, m)
	for i := 0; i < m; i++ {
		costGrid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			costGrid[i][j] = math.MaxInt
		}
	}

	costGrid[0][0] = 0
	var bfs func(i, j int)
	bfs = func(i, j int) {
		for k, move := range moves {
			newCost := costGrid[i][j]
			if grid[i][j] != k+1 {
				newCost++
			}
			x, y := i+move[0], j+move[1]
			if i >= 0 && i < m && j > 0 && j < n && newCost < costGrid[x][y] {
				costGrid[x][y] = newCost
				bfs(x, y)
			}
		}
	}
	bfs(0, 0)
	return costGrid[m-1][n-1]

}

// @lc code=end
