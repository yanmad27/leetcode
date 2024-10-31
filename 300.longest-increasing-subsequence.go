package main

import "sort"

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {

	tmp := []int{}
	for _, num := range nums {
		index := sort.SearchInts(tmp, num)
		if index == len(tmp) {
			tmp = append(tmp, num)
		} else {
			tmp[index] = num
		}
	}
	return len(tmp)
}

// @lc code=end
