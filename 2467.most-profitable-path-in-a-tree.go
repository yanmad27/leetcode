package main

import (
	"fmt"
	"slices"
)

/*
 * @lc app=leetcode id=2467 lang=golang
 *
 * [2467] Most Profitable Path in a Tree
 */

// @lc code=start
func mostProfitablePath(edges [][]int, bob int, amount []int) int {

	treeMap := make(map[int][]int)
	for _, edge := range edges {
		treeMap[edge[0]] = append(treeMap[edge[0]], edge[1])
		treeMap[edge[1]] = append(treeMap[edge[1]], edge[0])
	}

	var bobPath []int
	var bobGo func(cur int, path []int)
	bobGo = func(cur int, path []int) {
		path = append(path, cur)
		nodes := treeMap[cur]
		if cur == 0 {
			bobPath = path
			return
		}
		for _, node := range nodes {
			if slices.Contains(path, node) {
				continue
			}
			bobGo(node, path)
		}
	}

	fmt.Println(bobPath)
	return -1

}

// @lc code=end
