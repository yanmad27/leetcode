package main

import "strings"

/*
 * @lc app=leetcode id=1957 lang=golang
 *
 * [1957] Delete Characters to Make Fancy String
 */

// @lc code=start
func makeFancyString(s string) string {

	if len(s) < 3 {
		return s
	}
	b := strings.Builder{}
	b.WriteByte(s[0])
	b.WriteByte(s[1])
	for i := 2; i < len(s); i++ {
		if s[i] != s[i-1] || s[i] != s[i-2] {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

// @lc code=end
