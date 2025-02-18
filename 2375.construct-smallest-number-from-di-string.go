package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=2375 lang=golang
 *
 * [2375] Construct Smallest Number From DI String
 */

// @lc code=start
func usedI(num, i int) bool {
	for num > 0 {
		tmp := num % 10
		if tmp == i {
			return true
		}
		num = num / 10
	}
	return false
}
func smallestNumber(pattern string) string {
	n := len(pattern)
	result := math.MaxInt
	var recursive func(cur, len int)
	recursive = func(cur, len int) {
		if len == n {
			result = min(result, cur)
			return
		}
		lastNum := cur % 10
		for i := 1; i <= 9; i++ {
			if usedI(cur, i) {
				continue
			}
			if (pattern[len] == 'I' && i > lastNum) ||
				(pattern[len] == 'D' && i < lastNum) {
				recursive(cur*10+i, len+1)
			}
		}
	}

	for i := 1; i <= 9; i++ {
		recursive(i, 0)
	}
	return fmt.Sprintf("%d", result)
}

// @lc code=end
