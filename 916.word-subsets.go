package main

/*
 * @lc app=leetcode id=916 lang=golang
 *
 * [916] Word Subsets
 */

// @lc code=start
func isSubset(a, b string) bool {
	cnt := make([]int, 26)
	for _, c := range b {
		cnt[c-'a']++
	}
	for _, c := range a {
		cnt[c-'a']--
	}
	for _, v := range cnt {
		if v > 0 {
			return false
		}
	}
	return true
}
func wordSubsets(words1 []string, words2 []string) []string {

	result := []string{}
	for _, word1 := range words1 {
		isValid := true
		for _, word2 := range words2 {
			if !isSubset(word1, word2) {
				isValid = false
				break
			}
		}
		if isValid {
			result = append(result, word1)
		}
	}
	return result
}

// @lc code=end
