package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=3163 lang=golang
 *
 * [3163] String Compression III
 */

// @lc code=start
func compressedString(word string) string {

	word += "-"
	char := word[0]
	count := 1
	comp := strings.Builder{}
	for i := 1; i < len(word); i++ {
		if char == word[i] {
			count++
		} else {
			div := count / 9
			mod := count % 9
			for i := 0; i < div; i++ {
				comp.WriteString(fmt.Sprintf("%d%s", 9, string(char)))
			}
			if mod != 0 {
				comp.WriteString(fmt.Sprintf("%d%s", mod, string(char)))
			}
			char = word[i]
			count = 1
		}
	}
	return comp.String()
}

// @lc code=end
