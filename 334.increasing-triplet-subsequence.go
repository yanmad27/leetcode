package main

import "math"

/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))

// @lc code=start
func increasingTriplet(nums []int) bool {
	max1 := math.MaxInt
	max2 := math.MaxInt
	for _, num := range nums {
		if num > max2 {
			return true
		}
		if num > max1 {
			max2 = num
		} else {
			max1 = num
		}
	}
	return false
}

// @lc code=end
