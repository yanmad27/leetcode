package main

/*
 * @lc app=leetcode id=2425 lang=golang
 *
 * [2425] Bitwise XOR of All Pairings
 */

// @lc code=start
func xorAllNums(nums1 []int, nums2 []int) int {
	rs := 0
	if len(nums1)%2 == 1 {
		for _, v := range nums2 {
			rs ^= v
		}
	}
	if len(nums2)%2 == 1 {
		for _, v := range nums1 {
			rs ^= v
		}
	}
	return rs
}

// @lc code=end
