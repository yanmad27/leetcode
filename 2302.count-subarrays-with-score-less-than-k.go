package main

/*
 * @lc app=leetcode id=2302 lang=golang
 *
 * [2302] Count Subarrays With Score Less Than K
 */

// @lc code=start
func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	var sum, count int64
	for i, j := 0, 0; j < n; j++ {
		sum += int64(nums[j])
		for i <= j && sum*int64(j-i+1) >= k {
			sum -= int64(nums[i])
			i++
		}
		count += int64(j - i + 1)
	}
	return count
}

// @lc code=end
