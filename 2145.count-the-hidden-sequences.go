package main

/*
 * @lc app=leetcode id=2145 lang=golang
 *
 * [2145] Count the Hidden Sequences
 */

// @lc code=start
func numberOfArrays(differences []int, lower int, upper int) int {
	maxDiff := differences[0]
	minDiff := differences[0]
	cur := 0
	for _, num := range differences {
		cur += num
		maxDiff = max(maxDiff, cur)
		minDiff = min(minDiff, cur)
	}
	return max((upper-lower)-(maxDiff-minDiff)+1, 0)
}

// @lc code=end
