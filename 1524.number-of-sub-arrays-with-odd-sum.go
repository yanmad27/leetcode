package main

/*
 * @lc app=leetcode id=1524 lang=golang
 *
 * [1524] Number of Sub-arrays With Odd Sum
 */

// @lc code=start
func numOfSubarrays(arr []int) int {
	mod := 1e9 + 7
	odd, even, count, prefixSum := 0, 1, 0, 0
	for _, num := range arr {
		prefixSum += num
		if prefixSum%2 == 0 {
			count += odd
			even++
		} else {
			count += even
			odd++
		}
	}
	return count % int(mod)
}

// @lc code=end
