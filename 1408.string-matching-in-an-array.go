package main

import "strings"

/*
 * @lc app=leetcode id=1408 lang=golang
 *
 * [1408] String Matching in an Array
 */

// @lc code=start
func stringMatching(words []string) []string {
	result := []string{}
	for _, word1 := range words {
		for _, word2 := range words {
			if word1 != word2 && strings.Contains(word2, word1) {
				result = append(result, word1)
				break
			}
		}
	}
	return result
}

// @lc code=end
