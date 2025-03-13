package main

import "sort"

/*
 * @lc app=leetcode id=3356 lang=golang
 *
 * [3356] Zero Array Transformation II
 */

// @lc code=start
func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	var checkValid func(k int) bool
	checkValid = func(k int) bool {
		diffMap := make(map[int]int, n+1)
		for i := 0; i < k; i++ {
			from, to, val := queries[i][0], queries[i][1], queries[i][2]
			diffMap[from] += val
			diffMap[to+1] -= val
		}
		tmp := 0
		for i, num := range nums {
			tmp += diffMap[i]
			if num > tmp {
				return false
			}
		}
		return true
	}

	if !checkValid(len(queries)) {
		return -1
	}

	return sort.Search(len(queries), checkValid)
}

// @lc code=end
