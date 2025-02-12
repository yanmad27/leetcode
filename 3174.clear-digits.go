package main

/*
 * @lc app=leetcode id=3174 lang=golang
 *
 * [3174] Clear Digits
 */

// @lc code=start
func clearDigits(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			result = result[:len(result)-1]
		} else {
			result = result + string(ch)
		}
	}
	return result
}

// @lc code=end
