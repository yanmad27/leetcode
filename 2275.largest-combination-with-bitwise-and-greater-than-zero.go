package main

/*
 * @lc app=leetcode id=2275 lang=golang
 *
 * [2275] Largest Combination With Bitwise AND Greater Than Zero
 */

// @lc code=start
func largestCombination(candidates []int) int {

	maxLength := 0
	for bit := 0; bit < 24; bit++ {
		count := 0
		for _, num := range candidates {
			if num&(1<<bit) != 0 {
				count++
			}
		}
		if count > maxLength {
			maxLength = count
		}
	}
	return maxLength
}

// @lc code=end
