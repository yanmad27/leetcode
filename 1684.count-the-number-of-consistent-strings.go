package main

/*
 * @lc app=leetcode id=1684 lang=golang
 *
 * [1684] Count the Number of Consistent Strings
 */

// @lc code=start
func countConsistentStrings(allowed string, words []string) int {

	allowedMap := make(map[rune]bool)
	for _, char := range allowed {
		allowedMap[char] = true
	}

	count := 0
	for _, word := range words {
		tmp := true
		for _, char := range word {
			if !allowedMap[char] {
				tmp = false
				break
			}
		}
		if tmp {
			count++
		}
	}
	return count
}

// @lc code=end
