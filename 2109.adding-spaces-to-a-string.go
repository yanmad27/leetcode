package main

import "strings"

/*
 * @lc app=leetcode id=2109 lang=golang
 *
 * [2109] Adding Spaces to a String
 */

// @lc code=start
func addSpaces(s string, spaces []int) string {

	builder := strings.Builder{}
	spaceI := 0
	n := len(spaces)
	for i, c := range s {
		if spaceI < n && i == spaces[spaceI] {
			builder.WriteRune(' ')
			spaceI++
		}
		builder.WriteRune(c)
	}
	return builder.String()
}

// @lc code=end
