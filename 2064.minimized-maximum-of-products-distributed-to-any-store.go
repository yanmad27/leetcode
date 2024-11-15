package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=2064 lang=golang
 *
 * [2064] Minimized Maximum of Products Distributed to Any Store
 */

// @lc code=start
func minimizedMaximum(n int, quantities []int) int {

	canDistribute := func(k int) bool {
		sum := 0
		if k == 0 {
			return false
		}
		for _, q := range quantities {
			sum += q / k
			if q%k != 0 {
				sum++
			}
		}
		return sum <= n
	}
	return sort.Search(100000, func(i int) bool {
		return canDistribute(i)
	})
}

// @lc code=end
