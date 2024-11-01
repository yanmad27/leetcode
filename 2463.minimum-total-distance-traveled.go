package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=2463 lang=golang
 *
 * [2463] Minimum Total Distance Traveled
 */

// @lc code=start
func abs(value int) int {
	if value > 0 {
		return value
	}
	return -value
}
func minimumTotalDistance(robots []int, factories [][]int) int64 {
	sort.Slice(robots, func(i, j int) bool {
		return robots[i] < robots[j]
	})
	sort.Slice(factories, func(i, j int) bool {
		return factories[i][0] < factories[j][0]
	})
	minDistance := math.MaxInt
	var support func(robot []int, factory [][]int, distance int)
	support = func(robot []int, factory [][]int, distance int) {
		for i := 0; i < len(factory); i++ {
			if factory[i][1] > 0 {
				tmp := abs(robot[0] - factory[i][0])
				factory[i][1]--
				newDistance := distance + tmp
				if len(robot) == 1 {
					minDistance = min(minDistance, newDistance)
				} else if newDistance < minDistance {
					support(robot[1:], factory, newDistance)
				}
				factory[i][1]++
			}
		}
	}

	support(robots, factories, 0)
	return int64(minDistance)
}

// @lc code=end
