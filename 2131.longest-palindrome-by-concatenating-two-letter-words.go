package main

/*
 * @lc app=leetcode id=2131 lang=golang
 *
 * [2131] Longest Palindrome by Concatenating Two Letter Words
 */

// @lc code=start
func longestPalindrome(words []string) int {
	countMap := make(map[string]int)
	for _, word := range words {
		countMap[word]++
	}

	maxLength := 0
	selected := false
	for key, val := range countMap {
		if val == 0 {
			continue
		}
		if key[0] != key[1] {
			reverseKey := string(key[1]) + string(key[0])
			reverseValue := countMap[reverseKey]
			maxLength += min(val, reverseValue) * 4
			countMap[key] = 0
			countMap[reverseKey] = 0
		} else {
			if val%2 == 0 {
				maxLength += val * 2
			} else {
				if !selected {
					maxLength += val * 2
					selected = true
				} else {
					maxLength += (val - 1) * 2
				}
			}

		}
	}
	return maxLength

}

// @lc code=end
