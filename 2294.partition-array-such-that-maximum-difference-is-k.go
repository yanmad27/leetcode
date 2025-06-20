package main

import "sort"

/*
 * @lc app=leetcode id=2294 lang=golang
 *
 * [2294] Partition Array Such That Maximum Difference Is K
 */

// @lc code=start
func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	count := 1
	first := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num-first <= k {
			continue
		}
		first = num
		count++
	}
	return count
}

// @lc code=end
