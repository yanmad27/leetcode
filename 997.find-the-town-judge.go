package main

/*
 * @lc app=leetcode id=997 lang=golang
 *
 * [997] Find the Town Judge
 */

// @lc code=start
func findJudge(n int, trust [][]int) int {
	trustMap := map[int]int{}
	for _, t := range trust {
		trustMap[t[1]]++
	}
	maxCount := 0
	townJudge := 1
	for body, trustCount := range trustMap {
		if trustCount > maxCount {
			maxCount = trustCount
			townJudge = body
		}
	}
	if maxCount != n-1 {
		return -1
	}
	for _, t := range trust {
		if t[0] == townJudge {
			return -1
		}
	}
	return townJudge
}

// @lc code=end
