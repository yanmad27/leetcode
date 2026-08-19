package main

/*
 * @lc app=leetcode id=1679 lang=golang
 *
 * [1679] Max Number of K-Sum Pairs
 */

// @lc code=start
func maxOperations(nums []int, k int) int {
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num]++
	}
	count := 0
	for _, num := range nums {
		if k-num == num {
			count += maps[num] / 2
			maps[num] = 0
		} else {
			if _, ok := maps[k-num]; ok {
				count += min(maps[k-num], maps[num])
				maps[k-num] = 0
				maps[num] = 0
			}
		}
	}
	return count
}

// @lc code=end
