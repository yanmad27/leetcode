package main

import "slices"

/*
 * @lc app=leetcode id=1079 lang=golang
 *
 * [1079] Letter Tile Possibilities
 */

// @lc code=start
func numTilePossibilities(tiles string) int {
	resultMap := make(map[string]bool)
	n := len(tiles)
	var recursive func(cur []byte, used []int, i, length int)
	recursive = func(cur []byte, used []int, i, length int) {
		if length > n {
			return
		}
		if slices.Contains(used, i) {
			return
		}
		cur = append(cur, tiles[i])
		resultMap[string(cur)] = true
		used = append(used, i)
		for new := 0; new < n; new++ {
			recursive(cur, used, new, length+1)
		}
	}
	for i := 0; i < n; i++ {
		recursive([]byte{}, []int{}, i, 0)
	}
	return len(resultMap)
}

// @lc code=end
