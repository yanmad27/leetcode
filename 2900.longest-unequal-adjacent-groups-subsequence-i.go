package main

/*
 * @lc app=leetcode id=2900 lang=golang
 *
 * [2900] Longest Unequal Adjacent Groups Subsequence I
 */

// @lc code=start
func getLongestSubsequence(words []string, groups []int) []string {
	result := []string{words[0]}
	for i := 1; i < len(groups); i++ {
		if groups[i] != groups[i-1] {
			result = append(result, words[i])
		}
	}
	return result
}

// @lc code=end
