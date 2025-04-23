package main

/*
 * @lc app=leetcode id=1399 lang=golang
 *
 * [1399] Count Largest Group
 */

// @lc code=start
func countLargestGroup(n int) int {
	maxCount := 0
	count := make(map[int]int)
	for i := 1; i <= n; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		count[sum]++
		maxCount = max(maxCount, count[sum])
	}
	result := 0
	for _, v := range count {
		if v == maxCount {
			result++
		}
	}
	return result
}

// @lc code=end
