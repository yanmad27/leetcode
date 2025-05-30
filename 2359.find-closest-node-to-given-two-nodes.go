package main

import "math"

/*
 * @lc app=leetcode id=2359 lang=golang
 *
 * [2359] Find Closest Node to Given Two Nodes
 */

// @lc code=start
func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	var getDistance func(int) []int
	getDistance = func(node int) []int {
		distance := make([]int, n)
		visited := make([]bool, n)
		for i := 0; i < n; i++ {
			distance[i] = -1
		}
		count := 0
		for node != -1 && !visited[node] {
			distance[node] = count
			visited[node] = true
			count++
			node = edges[node]
		}
		return distance
	}
	dfn1 := getDistance(node1)
	dfn2 := getDistance(node2)

	tmp1 := []int{}
	ind1 := []int{}
	for i := 0; i < n; i++ {
		if dfn1[i] == -1 || dfn2[i] == -1 {
			continue
		}
		tmp1 = append(tmp1, max(dfn1[i], dfn2[i]))
		ind1 = append(ind1, i)
	}
	minNode := -1
	minDistance := math.MaxInt
	for i := 0; i < len(tmp1); i++ {
		if tmp1[i] < minDistance {
			minDistance = tmp1[i]
			minNode = ind1[i]
		} else if tmp1[i] == minDistance {
			minNode = min(minNode, ind1[i])
		}
	}

	return minNode

}

// @lc code=end
