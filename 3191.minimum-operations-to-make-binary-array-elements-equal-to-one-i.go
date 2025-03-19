package main

/*
 * @lc app=leetcode id=3191 lang=golang
 *
 * [3191] Minimum Operations to Make Binary Array Elements Equal to One I
 */

// @lc code=start
func minOperations(nums []int) int {
	n := len(nums)
	operator := 0
	for i := 0; i <= n-3; i++ {
		if nums[i] == 0 {
			operator++
			nums[i+1] = 1 - nums[i+1]
			nums[i+2] = 1 - nums[i+2]
		}
	}
	if nums[len(nums)-1] == 1 && nums[len(nums)-2] == 1 {
		return operator
	}
	return -1
}

// @lc code=end
