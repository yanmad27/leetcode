package main

/*
 * @lc app=leetcode id=1780 lang=golang
 *
 * [1780] Check if Number is a Sum of Powers of Three
 */

// @lc code=start
func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 > 1 {
			return false
		}
		n /= 3
	}
	return true
}

// @lc code=end
