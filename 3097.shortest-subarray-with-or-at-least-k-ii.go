package main

import "math"

/*
 * @lc app=leetcode id=3097 lang=golang
 *
 * [3097] Shortest Subarray With OR at Least K II
 */

// @lc code=start

func minimumSubarrayLength(nums []int, k int) int {

	result := math.MaxInt
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			return 1
		}
		n = n | nums[i]
		if n >= k {
			t := nums[i]
			j := i
			for t < k {
				j--
				t = t | nums[j]
			}
			if i-j+1 < result {
				result = i - j + 1
			}
			j++
			n = nums[j]
			i = j
		}
	}
	if result < math.MaxInt {
		return result
	}
	return -1

}

// @lc code=end
