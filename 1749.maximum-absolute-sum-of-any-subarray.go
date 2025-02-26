package main

/*
 * @lc app=leetcode id=1749 lang=golang
 *
 * [1749] Maximum Absolute Sum of Any Subarray
 */

// @lc code=start

func maxAbsoluteSum(nums []int) int {
	maxSum := 0
	sum1 := 0
	sum2 := 0
	for _, num := range nums {
		sum1 += num
		sum2 -= num
		if sum1 < 0 {
			sum1 = 0
		} else if sum1 > maxSum {
			maxSum = sum1
		}

		if sum2 < 0 {
			sum2 = 0
		} else if sum2 > maxSum {
			maxSum = sum2
		}
	}
	return maxSum
}

// @lc code=end
