package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=670 lang=golang
 *
 * [670] Maximum Swap
 */

// @lc code=start
func maximumSwap(num int) int {

	str := []rune(fmt.Sprintf("%d", num))

	maximum := num
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] < str[j] {
				str[i], str[j] = str[j], str[i]
				var tmp int
				fmt.Sscanf(string(str), "%d", &tmp)
				maximum = max(maximum, tmp)
				str[i], str[j] = str[j], str[i]
			}
		}
	}
	return maximum
}

// @lc code=end
