package main

/*
 * @lc app=leetcode id=2140 lang=golang
 *
 * [2140] Solving Questions With Brainpower
 */

// @lc code=start
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		points := questions[i][0]
		skip := questions[i][1]
		nextPos := min(n, i+skip+1)
		dp[i] = max(points+dp[nextPos], dp[i+1])
	}

	return int64(dp[0])
}

// @lc code=end
