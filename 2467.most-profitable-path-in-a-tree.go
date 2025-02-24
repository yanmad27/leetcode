package main

import (
	"math"
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

	bobPath := make(map[int]int)
	bobMove := make(map[int]int)
	foundBobPath := false
	var findBobPath func(cur, time int)
	findBobPath = func(cur, time int) {
		if foundBobPath {
			return
		}
		bobPath[cur] = time
		bobMove[time] = cur
		if cur == 0 {
			foundBobPath = true
			return
		}
		for _, nextMove := range treeMap[cur] {
			if _, ok := bobPath[nextMove]; ok {
				continue
			}
			findBobPath(nextMove, time+1)
			if !foundBobPath {
				delete(bobPath, nextMove)
				delete(bobMove, time+1)
			}
		}
	}
	findBobPath(bob, 0)

	maxIncome := math.MinInt
	var aliceGo func(cur, from, income, time int)
	aliceGo = func(cur, from, income, time int) {
		if len(treeMap[cur]) == 1 && cur != 0 {
			maxIncome = max(maxIncome, income)
			return
		}
		for _, node := range treeMap[cur] {
			if node == from {
				continue
			}
			newIncome := amount[node]
			if node == bobMove[time] {
				newIncome = amount[node] / 2
			} else if t, ok := bobPath[node]; ok && time > t {
				newIncome = 0
			}
			aliceGo(node, cur, income+newIncome, time+1)
		}
	}
	aliceGo(0, -1, amount[0], 1)

	return maxIncome

}

// @lc code=end
