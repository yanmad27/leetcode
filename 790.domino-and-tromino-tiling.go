package main

/*
 * @lc app=leetcode id=790 lang=golang
 *
 * [790] Domino and Tromino Tiling
 */

// @lc code=start
func numTilings(n int) int {
	dp := make([]int, n+4)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % 1000000007
	}
	return dp[n]
}

// @lc code=end
