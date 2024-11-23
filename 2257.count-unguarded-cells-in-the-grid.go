package main

/*
 * @lc app=leetcode id=2257 lang=golang
 *
 * [2257] Count Unguarded Cells in the Grid
 */

// @lc code=start
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {

	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	for _, wall := range walls {
		grid[wall[0]][wall[1]] = 1
	}
	for _, guard := range guards {
		x, y := guard[0], guard[1]
		grid[x][y] = 1
		// left
		for j := y - 1; j >= 0; j-- {
			if grid[x][j] == 0 {
				grid[x][j] = 2
			} else if grid[x][j] == 1 || grid[x][j] == 2 {
				break
			}
		}

		// right
		for j := y + 1; j < n; j++ {
			if grid[x][j] == 0 {
				grid[x][j] = 2
			} else if grid[x][j] == 1 || grid[x][j] == 2 {
				break
			}
		}

		// top
		for i := x - 1; i >= 0; i-- {
			if grid[i][y] == 0 {
				grid[i][y] = 3
			} else if grid[i][y] == 1 || grid[i][y] == 3 {
				break
			}
		}

		// bottom
		for i := x + 1; i < m; i++ {
			if grid[i][y] == 0 {
				grid[i][y] = 3
			} else if grid[i][y] == 1 || grid[i][y] == 3 {
				break
			}
		}

	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
			}
		}
	}
	return count

}

// @lc code=end
