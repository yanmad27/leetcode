package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	split := strings.Split(s, " ")
	rs := make([]string, 0, len(split))
	for i := len(split) - 1; i >= 0; i-- {
		if split[i] != "" {
			rs = append(rs, split[i])
		}
	}

	return strings.Join(rs, " ")
}

// @lc code=end
