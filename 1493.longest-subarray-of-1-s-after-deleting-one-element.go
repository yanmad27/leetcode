package main

/*
 * @lc app=leetcode id=1493 lang=golang
 *
 * [1493] Longest Subarray of 1's After Deleting One Element
 */

// @lc code=start
func longestSubarray(nums []int) int {

	n := len(nums)
	left, right := 0, 0
	countZero := 0
	result := 0

	for right < n {
		if nums[right] == 0 {
			countZero++
		}
		if countZero > 1 {
			for countZero > 1 {
				if nums[left] == 0 {
					countZero--
				}
				left++
			}
		} else {
			result = max(result, right-left+1)
		}
		right++
	}
	return result-1
}

// @lc code=end
