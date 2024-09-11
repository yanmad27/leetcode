package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	length := 0
	maxLength := 0
	existedMap := make(map[byte]bool)
	for r < len(s) {
		if existed := existedMap[s[r]]; existed {
			for s[l] != s[r] {
				existedMap[s[l]] = false
				l++
				length--
			}
			l++
			length--
		}
		existedMap[s[r]] = true
		r++
		length++
		if length > maxLength {
			maxLength = length
		}

	}

	return maxLength
}

// @lc code=end
