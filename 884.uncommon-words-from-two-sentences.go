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
	s := s1 + " " + s2
	arr := strings.Split(s, " ")
	for _, word := range arr {
		_, ok := commonMap[word]
		commonMap[word] = !ok
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
