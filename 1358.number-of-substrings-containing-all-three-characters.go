package main

/*
 * @lc app=leetcode id=1358 lang=golang
 *
 * [1358] Number of Substrings Containing All Three Characters
 */

// @lc code=start
func numberOfSubstrings(s string) int {
	countMap := make(map[rune]int, 3)
	countMap['a'], countMap['b'], countMap['c'] = -1, -1, -1
	result := 0
	for i, c := range s {
		countMap[c] = i

		if countMap['a'] >= 0 && countMap['b'] >= 0 && countMap['c'] >= 0 {
			minIndex := min(countMap['a'], countMap['c'], countMap['b'])
			result += minIndex + 1
		}
	}
	return result
}

// @lc code=end
