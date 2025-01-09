package main

import "strings"

/*
 * @lc app=leetcode id=2185 lang=golang
 *
 * [2185] Counting Words With a Given Prefix
 */

// @lc code=start
func prefixCount(words []string, pref string) int {
	count := 0

	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			count++
		}
	}
	return count
}

// @lc code=end
