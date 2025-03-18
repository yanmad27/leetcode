package main

/*
 * @lc app=leetcode id=2401 lang=golang
 *
 * [2401] Longest Nice Subarray
 */

// @lc code=start
func longestNiceSubarray(nums []int) int {
	n := len(nums)
	l := 0
	result := 1
	for r := 1; r < n; r++ {
		isNice := true
		for i := l; i < r; i++ {
			if nums[r]&nums[i] != 0 {
				isNice = false
				break
			}
		}
		if isNice {
			result = max(result, r-l+1)
		} else {
			l++
		}
	}
	return result
}

// @lc code=end
