package main

/*
 * @lc app=leetcode id=1769 lang=golang
 *
 * [1769] Minimum Number of Operations to Move All Balls to Each Box
 */

// @lc code=start
func lc1769_minOperations(boxes string) []int {
	hasBoxIndex := []int{}
	for i := 0; i < len(boxes); i++ {
		if boxes[i] == '1' {
			hasBoxIndex = append(hasBoxIndex, i)
		}
	}
	res := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		for _, index := range hasBoxIndex {
			res[i] += abs(index - i)
		}
	}
	return res
}

// @lc code=end
