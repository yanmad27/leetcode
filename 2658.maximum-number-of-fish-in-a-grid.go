package main

/*
 * @lc app=leetcode id=2658 lang=golang
 *
 * [2658] Maximum Number of Fish in a Grid
 */

// @lc code=start
func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 0
		}
		fish := grid[x][y]
		grid[x][y] = 0
		return fish +
			dfs(x-1, y) +
			dfs(x+1, y) +
			dfs(x, y-1) +
			dfs(x, y+1)
	}
	maxFish := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			maxFish = max(maxFish, dfs(i, j))
		}
	}
	return maxFish
}

// @lc code=end
