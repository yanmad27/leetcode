package main

import "fmt"

/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {

	rs := 0
	prevChar := chars[0]
	cnt := 1
	for i := 1; i < len(chars); i++ {
		char := chars[i]
		if char == prevChar {
			cnt++
			continue
		}
		chars[rs] = prevChar
		rs++
		if cnt != 1 {
			s := fmt.Sprintf("%d", cnt)
			for k := 0; k < len(s); k++ {
				chars[rs] = s[k]
				rs++
			}
		}
		prevChar = char
		cnt = 1
	}
	chars[rs] = prevChar
	rs++
	if cnt != 1 {
		s := fmt.Sprintf("%d", cnt)
		for k := 0; k < len(s); k++ {
			chars[rs] = s[k]
			rs++
		}
	}
	return rs
}

// @lc code=end
