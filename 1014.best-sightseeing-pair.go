package main

/*
 * @lc app=leetcode id=1014 lang=golang
 *
 * [1014] Best Sightseeing Pair
 */

// @lc code=start
func maxScoreSightseeingPair(values []int) int {

	rs := values[0] - 1
	prevMax := values[0] - 1
	for i := 1; i < len(values); i++ {
		rs = max(rs, prevMax+values[i])
		prevMax = max(prevMax, values[i]) - 1
	}

	return rs
}

// @lc code=end
