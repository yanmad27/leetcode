package main

import "sort"

/*
 * @lc app=leetcode id=2779 lang=golang
 *
 * [2779] Maximum Beauty of an Array After Applying Operation
 */

// @lc code=start
func maximumBeauty(nums []int, k int) int {

	sort.Ints(nums)
	k = 2 * k
	i := 0
	j := 0
	l := len(nums)
	r := 0
	for j < l {
		for j < l && nums[i]+k >= nums[j] {
			if 1+j-i > r {
				r = j - i + 1
			}
			j++
		}
		i++
	}
	return r
}

// @lc code=end

// 1 2 4 6
