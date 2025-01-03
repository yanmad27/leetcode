package main

/*
 * @lc app=leetcode id=2270 lang=golang
 *
 * [2270] Number of Ways to Split Array
 */

// @lc code=start
func waysToSplitArray(nums []int) int {
	sumRight := 0
	for _, num := range nums {
		sumRight += num
	}
	sumLeft := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		sumLeft += nums[i]
		sumRight -= nums[i]
		if sumLeft >= sumRight {
			count++
		}
	}
	return count
}

// @lc code=end
