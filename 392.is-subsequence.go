package main

/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	is, it := 0, 0
	ls, lt := len(s), len(t)
	if ls == 0 {
		return true
	}
	for is < ls && it < lt {
		if s[is] == t[it] {
			is++
			it++
			if is == ls {
				return true
			}
		} else {
			it++
		}
	}
	return false
}

// @lc code=end
