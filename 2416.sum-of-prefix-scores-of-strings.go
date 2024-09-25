package main

/*
 * @lc app=leetcode id=2416 lang=golang
 *
 * [2416] Sum of Prefix Scores of Strings
 */

// @lc code=start
func sumPrefixScores(words []string) []int {
	result := make([]int, len(words))
	prefixMap := make(map[string]int)
	for _, word := range words {
		for i := 1; i <= len(word); i++ {
			prefix := word[:i]
			prefixMap[prefix] += 1
		}
	}
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			prefix := word[:j]
			result[i] += prefixMap[prefix]
		}
	}

	return result
}

// @lc code=end
