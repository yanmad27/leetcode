package main

/*
 * @lc app=leetcode id=2938 lang=golang
 *
 * [2938] Separate Black and White Balls
 */

// @lc code=start
func minimumSteps(s string) int64 {

	count := 0
	l, r := 0, len(s)-1
	for l < r {
		if s[l] == '0' {
			l++
			continue
		}
		if s[r] == '1' {
			r--
			continue
		}
		count += r - l
		r--
		l++
	}
	return int64(count)
}

// @lc code=end
