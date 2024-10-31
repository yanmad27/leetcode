package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=2501 lang=golang
 *
 * [2501] Longest Square Streak in an Array
 */

// @lc code=start
func longestSquareStreak(nums []int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	n := len(nums)
	dpMap := make(map[int]int, n)
	maxCount := 1

	for _, num := range nums {
		dpMap[num] = max(dpMap[num], dpMap[num*num]+1)
		maxCount = max(dpMap[num], maxCount)
	}
	if maxCount == 1 {
		return -1
	}
	return maxCount
}

// @lc code=end
