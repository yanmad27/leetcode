package main

import "strconv"

/*
 * @lc app=leetcode id=2566 lang=golang
 *
 * [2566] Maximum Difference by Remapping a Digit
 */

// @lc code=start
func minMaxDifference(num int) int {

	s := []byte(strconv.Itoa(num))
	var val byte
	for i := 0; i < len(s); i++ {
		if s[i] == '9' {
			continue
		}
		val = s[i]
		break
	}
	for i := 0; i < len(s); i++ {
		if s[i] == val {
			s[i] = '9'
		}
	}
	maxValue, _ := strconv.Atoi(string(s))
	s = []byte(strconv.Itoa(num))
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}
		val = s[i]
		break
	}
	for i := 0; i < len(s); i++ {
		if s[i] == val {
			s[i] = '0'
		}
	}
	minValue, _ := strconv.Atoi(string(s))

	return maxValue - minValue
}

// @lc code=end
