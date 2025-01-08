package main

import "strings"

/*
 * @lc app=leetcode id=3042 lang=golang
 *
 * [3042] Count Prefix and Suffix Pairs I
 */

// @lc code=start
func isPrefixAndSuffix(str1, str2 string) bool {
	return strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1)
}
func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

// @lc code=end
