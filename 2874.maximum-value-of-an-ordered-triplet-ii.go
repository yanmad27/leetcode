package main

/*
 * @lc app=leetcode id=2874 lang=golang
 *
 * [2874] Maximum Value of an Ordered Triplet II
 */

// @lc code=start
func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	tmp := nums[0]
	for i := 1; i < n-1; i++ {
		tmp = max(nums[i-1], tmp)
		prefixMax[i] = tmp
	}
	tmp = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		tmp = max(nums[i+1], tmp)
		suffixMax[i] = tmp
	}
	result := -1
	for i := 1; i <= n-2; i++ {
		result = max(result, (prefixMax[i]-nums[i])*suffixMax[i])
	}
	if result < 0 {
		return 0
	}
	return int64(result)
}

// @lc code=end
