package main

/*
 * @lc app=leetcode id=1671 lang=golang
 *
 * [1671] Minimum Number of Removals to Make Mountain Array
 */

// @lc code=start

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	left := make(map[int]int)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				left[i] = max(left[j]+1, left[i])
			}
		}
	}
	right := make(map[int]int)
	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] > nums[j] {
				right[i] = max(right[j]+1, right[i])
			}
		}
	}
	maxLength := 0
	for i := 1; i < n-1; i++ {
		if left[i] != 0 && right[i] != 0 {
			maxLength = max(maxLength, left[i]+right[i]+1)
		}
	}
	return n - maxLength
}

// @lc code=end
