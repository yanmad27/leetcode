package main

/*
 * @lc app=leetcode id=1920 lang=golang
 *
 * [1920] Build Array from Permutation
 */

// @lc code=start
func buildArray(nums []int) []int {

	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] += 1000 * (nums[nums[i]] % 1000)
	}
	for i := 0; i < n; i++ {
		nums[i] /= 1000
	}
	return nums
}

// @lc code=end
