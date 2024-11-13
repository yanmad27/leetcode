package main

import "sort"

/*
 * @lc app=leetcode id=2563 lang=golang
 *
 * [2563] Count the Number of Fair Pairs
 */

// @lc code=start
func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	support := func(value int) int64 {
		left := 0
		right := len(nums) - 1
		result := int64(0)
		for left < right {
			sum := nums[left] + nums[right]
			if sum < value {
				result += int64(right - left)
				left++
			} else {
				right--
			}
		}

		return result
	}
	return support(upper+1) - support(lower)
}

// @lc code=end
// 0 1 4 4 5 7
