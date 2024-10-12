package main

import "sort"

/*
 * @lc app=leetcode id=962 lang=golang
 *
 * [962] Maximum Width Ramp
 */

// @lc code=start
func maxWidthRamp(nums []int) int {

	pos := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		pos[i] = i
	}
	sort.Slice(pos, func(i, j int) bool {
		if nums[pos[i]] < nums[pos[j]] {
			return true
		}
		if nums[pos[i]] == nums[pos[j]] {
			return pos[i] < pos[j]
		}
		return false
	})
	maxRamp := 0
	tmp := 0
	for i := 1; i < len(pos); i++ {
		if pos[i] > pos[tmp] {
			maxRamp = max(maxRamp, pos[i]-pos[tmp])
		} else {
			tmp = i
		}
	}
	return maxRamp

}

// @lc code=end
