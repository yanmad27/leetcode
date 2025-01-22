package main

/*
 * @lc app=leetcode id=1765 lang=golang
 *
 * [1765] Map of Highest Peak
 */

// @lc code=start
func highestPeak(isWater [][]int) [][]int {
	moves := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m := len(isWater)
	n := len(isWater[0])
	queue := [][2]int{}
	result := make([][]int, m)
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		used[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				used[i][j] = true
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, move := range moves {
			i, j := cur[0], cur[1]
			x, y := i+move[0], j+move[1]
			if x >= 0 && x < m && y >= 0 && y < n && !used[x][y] {
				result[x][y] = result[i][j] + 1
				used[x][y] = true
				queue = append(queue, [2]int{x, y})
			}
		}
	}

	return result
}

// @lc code=end
