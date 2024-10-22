package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=1593 lang=golang
 *
 * [1593] Split a String Into the Max Number of Unique Substrings
 */

// @lc code=start

func copyMap(m map[string]bool) (c map[string]bool) {
	c = make(map[string]bool)
	for k, v := range m {
		c[k] = v
	}
	return c
}
func maxUniqueSplit(s string) int {

	arr := strings.Split(s, "")

	maximum := -1

	var recursive func(cur map[string]bool, sub string, index int)
	recursive = func(cur map[string]bool, sub string, index int) {
		if index >= len(arr) {
			maximum = max(len(cur), maximum)
			return
		}
		sub += arr[index]
		recursive(copyMap(cur), sub, index+1)
		cur2 := copyMap(cur)
		cur2[sub] = true
		sub = ""
		recursive(cur2, sub, index+1)
	}
	cur := make(map[string]bool)
	sub := ""
	recursive(cur, sub, 0)
	return maximum
}

// @lc code=end
