package main

import "strings"

/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) []string {
	commonMap := make(map[string]bool)
	s1s := strings.Split(s1, " ")
	s2s := strings.Split(s2, " ")
	for _, word := range s1s {
		if _, ok := commonMap[word]; !ok {
			commonMap[word] = true
		} else {
			commonMap[word] = false
		}
	}
	for _, word := range s2s {
		if _, ok := commonMap[word]; !ok {
			commonMap[word] = true
		} else {
			commonMap[word] = false
		}
	}
	result := []string{}
	for key, value := range commonMap {
		if value {
			result = append(result, key)
		}
	}
	return result
}

// @lc code=end
