package main

/*
 * @lc app=leetcode id=3223 lang=golang
 *
 * [3223] Minimum Length of String After Operations
 */

// @lc code=start
func minimumLength(s string) int {
	countMap := make(map[rune]int)

	for _, c := range s {
		countMap[c]++
	}
	result := 0
	for _, v := range countMap {
		if v <= 2 {
			result += v
		} else if v%2 == 0 {
			result += 2
		} else {
			result += 1
		}
	}
	return result
}

// @lc code=end
