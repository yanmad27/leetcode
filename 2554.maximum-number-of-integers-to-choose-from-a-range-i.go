package main

/*
 * @lc app=leetcode id=2554 lang=golang
 *
 * [2554] Maximum Number of Integers to Choose From a Range I
 */

// @lc code=start
func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := make(map[int]bool)
	for _, n := range banned {
		bannedMap[n] = true
	}
	tmpSum := 0
	count := 0
	for i := 1; i <= n; i++ {
		if !bannedMap[i] {
			tmpSum += i
			if tmpSum <= maxSum {
				count++
			} else {
				break
			}
		}
	}
	return count

}

// @lc code=end
