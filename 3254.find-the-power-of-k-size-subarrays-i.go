package main

/*
 * @lc app=leetcode id=3254 lang=golang
 *
 * [3254] Find the Power of K-Size Subarrays I
 */

// @lc code=start
func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	n := len(nums)
	rs := make([]int, n-k+1)
	for i := range len(rs) {
		rs[i] = -1
	}
	count := 1
	for i := 1; i < n; i++ {
		if nums[i]-1 == nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count == k {
			rs[i-k+1] = nums[i]
			count--
		}
	}
	return rs
}

// @lc code=end
