package main

/*
 * @lc app=leetcode id=3355 lang=golang
 *
 * [3355] Zero Array Transformation I
 */

// @lc code=start
func isZeroArray(nums []int, queries [][]int) bool {

	n := len(nums)
	decrease := make([]int, n+1)
	for _, query := range queries {
		decrease[query[0]] += -1
		decrease[query[1]+1] += 1
	}
	tmp := 0
	for i, num := range nums {
		tmp += decrease[i]
		num += tmp
		if num > 0 {
			return false
		}
	}
	return true
}

// @lc code=end
