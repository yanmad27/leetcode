package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=2616 lang=golang
 *
 * [2616] Minimize the Maximum Difference of Pairs
 */

// @lc code=start
func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	check := func(diff int) bool {
		cnt := 0
		for i := 0; i < n-1; i++ {
			if nums[i+1]-nums[i] < diff {
				cnt++
				i++
			}
			if cnt >= p {
				return true
			}
		}
		return false
	}

	return sort.Search(nums[n-1]-nums[0]+1, check) - 1
}

// @lc code=end
// 1 1 2 3 7 10

// 1 2 2 4
