package main

/*
 * @lc app=leetcode id=2981 lang=golang
 *
 * [2981] Find Longest Special Substring That Occurs Thrice I
 */

// @lc code=start
func maximumLength(s string) int {

	specials := make(map[string]int)
	tmp := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		specials[string(tmp)] = len(findAllIndex(s, string(tmp)))
		if s[i] != s[i-1] {
			tmp = []byte{s[i]}
		} else {
			tmp = append(tmp, s[i])
		}
	}
	specials[string(tmp)] = len(findAllIndex(s, string(tmp)))

	maxLength := -1
	for key, value := range specials {
		if value >= 3 && len(key) > maxLength {
			maxLength = len(key)
		}
	}

	return maxLength
}

// @lc code=end
