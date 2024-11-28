package main

import (
	"container/list"
	"math"
)

/*
 * @lc app=leetcode id=2290 lang=golang
 *
 * [2290] Minimum Obstacle Removal to Reach Corner
 */

// @lc code=start
func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := list.New()
	queue.PushBack([3]int{0, 0, 0})
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	minObstacles := math.MaxInt
	for queue.Len() > 0 {
		curr := queue.Front().Value.([3]int)
		queue.Remove(queue.Front())

		for _, dir := range directions {
			nextX := curr[0] + dir[0]
			nextY := curr[1] + dir[1]
			removed := curr[2]
			if nextX == m-1 && nextY == n-1 {
				minObstacles = min(minObstacles, curr[2])
				continue
			}
			if nextX >= 0 && nextY >= 0 && nextX < m && nextY < n && !visited[nextX][nextY] {
				visited[nextX][nextY] = true

				if grid[nextX][nextY] == 1 {
					removed++
					nextState := [3]int{nextX, nextY, removed}
					queue.PushBack(nextState)
				} else {
					nextState := [3]int{nextX, nextY, removed}
					queue.PushFront(nextState)
				}

			}
		}
	}
	return minObstacles
}

// @lc code=end
