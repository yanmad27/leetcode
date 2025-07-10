package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0
	for left < right {
		minHeight := min(height[left], height[right])
		result = max(result, minHeight*(right-left))

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return result
}

// @lc code=end
