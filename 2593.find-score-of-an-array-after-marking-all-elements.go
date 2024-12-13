package main

import "sort"

/*
 * @lc app=leetcode id=2593 lang=golang
 *
 * [2593] Find Score of an Array After Marking All Elements
 */

// @lc code=start
func findScore(nums []int) int64 {
	marked := make([]bool, len(nums))
	index := make([][2]int, len(nums))
	for i, num := range nums {
		index[i] = [2]int{num, i}
	}
	sort.Slice(index, func(i, j int) bool {
		if index[i][0] == index[j][0] {
			return index[i][1] < index[j][1]
		}
		return index[i][0] < index[j][0]
	})
	var score int64
	for _, i := range index {
		if marked[i[1]] {
			continue
		}
		score += int64(i[0])
		marked[i[1]] = true
		marked[max(i[1]-1, 0)] = true
		marked[min(i[1]+1, len(nums)-1)] = true
	}
	return score
}

// @lc code=end
