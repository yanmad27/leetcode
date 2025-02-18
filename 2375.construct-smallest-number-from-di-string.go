package main

import (
	"fmt"
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
	var recursive func(cur, len int) int
	recursive = func(cur, len int) int {
		if len == n {
			return cur
		}
		lastNum := cur % 10
		for i := 1; i <= 9; i++ {
			if usedI(cur, i) {
				continue
			}
			if (pattern[len] == 'I' && i > lastNum) ||
				(pattern[len] == 'D' && i < lastNum) {
				return recursive(cur*10+i, len+1)
			}
		}
		return -1
	}

	for i := 1; i <= 9; i++ {
		tmp := recursive(i, 0)
		if tmp != -1 {
			return fmt.Sprintf("%d", tmp)
		}

	}
	panic("unreachable")
}

// @lc code=end
