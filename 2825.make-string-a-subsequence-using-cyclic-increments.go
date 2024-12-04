package main

/*
 * @lc app=leetcode id=2825 lang=golang
 *
 * [2825] Make String a Subsequence Using Cyclic Increments
 */

// @lc code=start
func canMakeSubsequence(str1 string, str2 string) bool {
	l1, l2 := len(str1), len(str2)
	c1, c2 := 0, 0

	for c1 < l1 {
		tmp1 := str1[c1]
		tmp2 := str1[c1]
		if tmp2 == 'z' {
			tmp2 = 'a'
		} else {
			tmp2++
		}
		if tmp1 == str2[c2] || tmp2 == str2[c2] {
			c2++
			if c2 == l2 {
				return true
			}
		}
		c1++
	}
	return false

}

// @lc code=end
