package main

import "math"

/*
 * @lc app=leetcode id=3442 lang=golang
 *
 * [3442] Maximum Difference Between Even and Odd Frequency I
 */

// @lc code=start
func maxDifference(s string) int {
	arr := make([]int, 26)
	for _, char := range s {
		arr[char-'a']++
	}
	oddMax := 0
	evenMin := math.MaxInt
	for _, count := range arr {
		if count == 0 {
			continue
		}
		if count%2 == 0 && count < evenMin {
			evenMin = count
		} else if count%2 == 1 && count > oddMax {
			oddMax = count
		}
	}
	return oddMax - evenMin
}

// @lc code=end
