package main

/*
 * @lc app=leetcode id=3105 lang=golang
 *
 * [3105] Longest Strictly Increasing or Strictly Decreasing Subarray
 */

// @lc code=start
func longestMonotonicSubarray(nums []int) int {

	n := len(nums)
	maxIncrease := 1
	maxDecrease := 1
	for i := 0; i < n-1; i++ {
		increase := 1
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[j-1] {
				increase++
				maxIncrease = max(maxIncrease, increase)
			} else {
				break
			}
		}
		decrease := 1
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[j-1] {
				decrease++
				maxDecrease = max(maxDecrease, decrease)
			} else {
				break
			}
		}
	}

	return max(maxIncrease, maxDecrease)
}

// @lc code=end
