package main

import "strings"

/*
 * @lc app=leetcode id=2696 lang=golang
 *
 * [2696] Minimum String Length After Removing Substrings
 */

// @lc code=start
func minLength(s string) int {
	ab := strings.Index(s, "AB")
	cd := strings.Index(s, "CD")
	canContinue := ab != -1 || cd != -1
	for canContinue {
		if ab != -1 {
			s = strings.ReplaceAll(s, "AB", "")
		}
		if cd != -1 {
			s = strings.ReplaceAll(s, "CD", "")
		}
		ab = strings.Index(s, "AB")
		cd = strings.Index(s, "CD")
		canContinue = ab != -1 || cd != -1
	}
	return len(s)
}

// @lc code=end
