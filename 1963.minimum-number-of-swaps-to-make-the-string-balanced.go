package main

/*
 * @lc app=leetcode id=1963 lang=golang
 *
 * [1963] Minimum Number of Swaps to Make the String Balanced
 */

// @lc code=start
func minSwaps(s string) int {
	var close, maxClose int

	for _, c := range s {
		if c == ']' {
			close++
			maxClose = max(maxClose, close)
		} else {
			close--
		}
	}

	return (maxClose + 1) / 2
}

// @lc code=end
