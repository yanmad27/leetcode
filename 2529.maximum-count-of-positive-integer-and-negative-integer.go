package main

import "sort"

/*
 * @lc app=leetcode id=2529 lang=golang
 *
 * [2529] Maximum Count of Positive Integer and Negative Integer
 */

// @lc code=start
func maximumCount(nums []int) int {
	sort.Ints(nums)
	i := sort.SearchInts(nums, 0)
	j := sort.SearchInts(nums, 1)
	return max(len(nums)-j, i)
}

// @lc code=end
