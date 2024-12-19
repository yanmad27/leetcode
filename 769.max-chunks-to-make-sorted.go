package main

/*
 * @lc app=leetcode id=769 lang=golang
 *
 * [769] Max Chunks To Make Sorted
 */

// @lc code=start
func maxChunksToSorted(arr []int) int {
	maxMap := make(map[int]int)
	count := 0
	maxMap[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		maxMap[i] = max(arr[i], maxMap[i-1])
	}
	for i := 0; i < len(arr); i++ {
		if i == maxMap[i] {
			count++
		}
	}
	return count
}

// @lc code=end
