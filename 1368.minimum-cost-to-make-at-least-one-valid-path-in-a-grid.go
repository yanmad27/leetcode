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
	cost := make([][]int, m)
	for i := 0; i < m; i++ {
		cost[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cost[i][j] = math.MaxInt
		}
	}

	cost[0][0] = 0
	stack := [][2]int{{0, 0}}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		i, j := cur[0], cur[1]
		for k, move := range moves {
			newCost := cost[i][j]
			if grid[i][j]-1 != k {
				newCost++
			}
			x, y := i+move[0], j+move[1]
			if x >= 0 && x < m && y >= 0 && y < n && newCost < cost[x][y] {
				stack = append(stack, [2]int{x, y})
				cost[x][y] = newCost
			}
		}
	}

	return cost[m-1][n-1]

}

// @lc code=end
