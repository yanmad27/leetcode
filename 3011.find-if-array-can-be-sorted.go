package main

import (
	"math/bits"
)

/*
 * @lc app=leetcode id=3011 lang=golang
 *
 * [3011] Find if Array Can Be Sorted
 */

// @lc code=start
func canSwap(x, y int) bool {
	return bits.OnesCount(uint(x)) == bits.OnesCount(uint(y))
}
func canSortArray(nums []int) bool {

	for j := 1; j < len(nums); j++ {
		for i := 0; i < len(nums)-j; i++ {
			if nums[i] > nums[i+1] && canSwap(nums[i], nums[i+1]) {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

// @lc code=end

// 75: 1001011
// 34: 100010
// 30: 11110
