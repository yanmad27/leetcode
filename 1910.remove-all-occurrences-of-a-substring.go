package main

/*
 * @lc app=leetcode id=1910 lang=golang
 *
 * [1910] Remove All Occurrences of a Substring
 */

// @lc code=start
func removeOccurrences(s string, part string) string {
	m, n := len(s), len(part)
	for i := 0; i < m-n+1; i++ {
		occurred := true
		for j := 0; j < n; j++ {
			if s[i+j] != part[j] {
				occurred = false
				break
			}
		}
		if occurred {
			s = s[:i] + s[i+n:]
			i = -1
			m -= n
		}
	}
	return s
}

// @lc code=end
