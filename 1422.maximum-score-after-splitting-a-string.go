package main

/*
 * @lc app=leetcode id=1422 lang=golang
 *
 * [1422] Maximum Score After Splitting a String
 */

// @lc code=start
func maxScore(s string) int {
	one := 0
	for _, c := range s {
		if c == '1' {
			one++
		}
	}
	zero := 0
	max := 0
	for i := 0; i < len(s)-1; i++ {
		c := s[i]
		if c == '0' {
			zero++
		} else {
			one--
		}
		if zero+one > max {
			max = zero + one
		}
	}
	return max
}

// @lc code=end
