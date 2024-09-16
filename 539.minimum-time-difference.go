package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=539 lang=golang
 *
 * [539] Minimum Time Difference
 */

// @lc code=start
func findMinDifference(timePoints []string) int {
	minutes := make(map[int]bool)
	min := 24 * 60
	max := 0
	for _, time := range timePoints {
		tmp := strings.Split(time, ":")
		h, _ := strconv.Atoi(tmp[0])
		m, _ := strconv.Atoi(tmp[1])

		minute := h*60 + m
		if minutes[minute] {
			return 0
		}
		minutes[minute] = true
		if minute > max {
			max = minute
		}
		if minute < min {
			min = minute
		}
	}
	max = 1440 + min
	minutes[max] = true
	diff := 1440

	for i := min + 1; i <= max; i++ {
		if minutes[i] {
			if i-min < diff {
				diff = i - min
			}
			min = i
		}
	}
	return diff
}

// @lc code=end
