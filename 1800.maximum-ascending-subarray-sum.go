package main

/*
 * @lc app=leetcode id=1800 lang=golang
 *
 * [1800] Maximum Ascending Subarray Sum
 */

// @lc code=start
func maxAscendingSum(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		tmp := nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[j-1] {
				tmp += nums[j]
			} else {
				break
			}
		}
		sum = max(sum, tmp)
	}
	return sum
}

// @lc code=end
