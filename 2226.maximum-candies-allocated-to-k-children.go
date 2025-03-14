package main

import (
	"slices"
	"sort"
)

/*
 * @lc app=leetcode id=2226 lang=golang
 *
 * [2226] Maximum Candies Allocated to K Children
 */

// @lc code=start
func maximumCandies(candies []int, k int64) int {
	maxx := slices.Max(candies)
	return sort.Search(maxx, func(x int) bool {
		x++
		sum := 0
		for _, candy := range candies {
			sum += candy / x
		}
		return sum < int(k)
	})
}

// @lc code=end
