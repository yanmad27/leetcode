package main

/*
 * @lc app=leetcode id=2924 lang=golang
 *
 * [2924] Find Champion II
 */

// @lc code=start
func findChampion(n int, edges [][]int) int {

	teamMap := make(map[int]bool)
	for i := 0; i < n; i++ {
		teamMap[i] = true
	}
	for _, edge := range edges {
		delete(teamMap, edge[1])
	}
	if len(teamMap) >= 2 {
		return -1
	}
	for team := range teamMap {
		return team
	}
	return -1
}

// @lc code=end
