package main

/*
 * @lc app=leetcode id=2914 lang=golang
 *
 * [2914] Minimum Number of Changes to Make Binary String Beautiful
 */

// @lc code=start
func minChanges(s string) int {

	result := 0
	previous := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == previous {
			count++
			continue
		}
		if count%2 == 0 {
			previous = s[i]
			count = 1
			continue
		}

		result++
		tmp := min(i+1, len(s)-1)
		previous = s[tmp]
		count = 1
		i++
	}
	return result
}

// @lc code=end
