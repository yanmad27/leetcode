package main

/*
 * @lc app=leetcode id=2461 lang=golang
 *
 * [2461] Maximum Sum of Distinct Subarrays With Length K
 */

// @lc code=start
func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	countMap := make(map[int]int)
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
		countMap[nums[i]]++
	}
	maxSum := 0
	if len(countMap) == k {
		maxSum = sum
	}
	for i := k; i < n; i++ {
		sum -= nums[i-k]
		countMap[nums[i-k]]--
		if countMap[nums[i-k]] == 0 {
			delete(countMap, nums[i-k])
		}
		sum += nums[i]
		countMap[nums[i]]++
		if len(countMap) == k {
			maxSum = max(maxSum, sum)
		}
	}
	return int64(maxSum)
}

// @lc code=end
