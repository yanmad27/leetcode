package main

/*
 * @lc app=leetcode id=2683 lang=golang
 *
 * [2683] Neighboring Bitwise XOR
 */

// @lc code=start
func doesValidArrayExist(derived []int) bool {
	rs := 0
	for _, v := range derived {
		rs ^= v
	}
	return rs == 0
}

// @lc code=end
