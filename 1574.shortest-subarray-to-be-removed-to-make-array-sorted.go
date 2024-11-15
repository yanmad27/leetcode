package main

import "sort"

/*
 * @lc app=leetcode id=1574 lang=golang
 *
 * [1574] Shortest Subarray to be Removed to Make Array Sorted
 */

// @lc code=start
func findLengthOfShortestSubarray(arr []int) int {

	if sort.IntsAreSorted(arr) {
		return 0
	}
	n, l, r := len(arr), 0, 0
	for l = 0; l < n-1; l++ {
		if arr[l] > arr[l+1] {
			break
		}
	}
	for r = n - 1; r >= 1; r-- {
		if arr[r-1] > arr[r] {
			break
		}
	}
	maxLength := max(l+1, n-r)
	for i := 0; i <= l; i++ {
		for j := r; j < n; j++ {
			if arr[i] <= arr[j] {
				maxLength = max(maxLength, i+n-j+1)
				break
			}
		}
	}

	return n - maxLength

}

// @lc code=end
