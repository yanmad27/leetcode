package main

/*
 * @lc app=leetcode id=2044 lang=golang
 *
 * [2044] Count Number of Maximum Bitwise-OR Subsets
 */

// @lc code=start
func countMaxOrSubsets(nums []int) int {
	max := 0
	count := 1
	var recursive func(orValue, num int)
	recursive = func(orValue, i int) {
		if i == len(nums) {
			if orValue > max {
				max = orValue
				count = 1
			} else if orValue == max {
				count++
			}
			return
		}
		recursive(orValue, i+1)
		recursive(orValue|nums[i], i+1)
	}
	recursive(0, 0)
	return count
}

// @lc code=end
