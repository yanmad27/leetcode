package main

/*
 * @lc app=leetcode id=1752 lang=golang
 *
 * [1752] Check if Array Is Sorted and Rotated
 */

// @lc code=start
func check(nums []int) bool {
	swap := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] > nums[(i+1)%n] {
			swap++
		}
	}
	return swap <= 1
}

// @lc code=end
