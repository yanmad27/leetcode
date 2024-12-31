package main

/*
 * @lc app=leetcode id=2466 lang=golang
 *
 * [2466] Count Ways To Build Good Strings
 */

// @lc code=start

func countGoodStrings(low int, high int, zero int, one int) int {
	mod := 1000000007
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] += dp[i-zero]
		}
		if i >= one {
			dp[i] += dp[i-one]
		}
		dp[i] %= mod
	}

	res := 0
	for i := low; i <= high; i++ {
		res += dp[i]%mod
	}
	return res % mod

}

// @lc code=end
