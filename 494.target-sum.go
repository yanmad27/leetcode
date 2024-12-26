package main

/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	count := 0
	var dfs func(int, int)
	dfs = func(i, sum int) {
		if i == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		dfs(i+1, sum-nums[i])
		dfs(i+1, sum+nums[i])
	}
	dfs(0, 0)
	return count
}

// @lc code=end
