package main

/*
 * @lc app=leetcode id=769 lang=golang
 *
 * [769] Max Chunks To Make Sorted
 */

// @lc code=start
func maxChunksToSorted(arr []int) int {
	mem := make(map[int]bool)
	count := 0
	for i, num := range arr {
		mem[num] = true
		filled := true
		for j := 0; j <= i; j++ {
			if !mem[j] {
				filled = false
				break
			}
		}
		if filled {
			count++
		}
	}
	return count
}

// @lc code=end
