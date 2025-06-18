package main

import "sort"

/*
 * @lc app=leetcode id=2966 lang=golang
 *
 * [2966] Divide Array Into Arrays With Max Difference
 */

// @lc code=start
func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums); i += 3 {
		rs := []int{}
		for j := i; j < i+3; j++ {
			if len(rs) == 0 {
				rs = append(rs, nums[j])
				continue
			}
			if nums[j]-rs[0] <= k {
				rs = append(rs, nums[j])
			} else {
				return [][]int{}
			}
		}
		result = append(result, rs)
	}
	return result
}

// @lc code=end
