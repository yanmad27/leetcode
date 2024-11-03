package main

import "strings"

/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	for i := 0; i < len(s); i++ {
		if s[i+1:]+s[:i+1] == goal {
			return true
		}
	}
	return false
}

func rotateString2(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}

// @lc code=end
