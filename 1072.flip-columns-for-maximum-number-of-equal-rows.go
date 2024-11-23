package main

import "strings"

/*
 * @lc app=leetcode id=1072 lang=golang
 *
 * [1072] Flip Columns For Maximum Number of Equal Rows
 */

// @lc code=start
func maxEqualRowsAfterFlips(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	keyMap := make(map[string]int)
	rs := 0
	for i := 0; i < m; i++ {
		var key strings.Builder
		valid := matrix[i][0] == 0
		if valid {
			for j := 0; j < n; j++ {
				key.WriteRune(rune('0' + matrix[i][j]))
			}
		} else {
			for j := 0; j < n; j++ {
				key.WriteRune(rune('1' - matrix[i][j]))
			}
		}
		strKey := key.String()
		keyMap[strKey]++
		if keyMap[strKey] > rs {
			rs = keyMap[strKey]
		}
	}
	return rs
}

// @lc code=end
