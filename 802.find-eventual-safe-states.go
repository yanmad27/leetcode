package main

/*
 * @lc app=leetcode id=802 lang=golang
 *
 * [802] Find Eventual Safe States
 */

// @lc code=start
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	ternimalNode := []int{}
	for i, gr := range graph {
		if len(gr) == 0 {
			ternimalNode = append(ternimalNode, i)
		}
	}

	return []int{n}
}

// @lc code=end
