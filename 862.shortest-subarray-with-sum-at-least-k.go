package main

/*
 * @lc app=leetcode id=862 lang=golang
 *
 * [862] Shortest Subarray with Sum at Least K
 */

// @lc code=start
func shortestSubarray(nums []int, k int) int {

	n := len(nums)
	sum := make([]int, n+1)
	for i, num := range nums {
		sum[i+1] = sum[i] + num
	}
	dp := []int{}
	rs := n + 1
	for i := 0; i <= n; i++ {
		for len(dp) > 0 && sum[i]-sum[dp[0]] >= k {
			if i-dp[0] < rs {
				rs = i - dp[0]
			}
			dp = dp[1:]
		}

		for len(dp) > 0 && sum[i] <= sum[dp[len(dp)-1]] {
			dp = dp[:len(dp)-1]
		}

		dp = append(dp, i)
	}

	if rs == n+1 {
		return -1
	}
	return rs
}

// @lc code=end
