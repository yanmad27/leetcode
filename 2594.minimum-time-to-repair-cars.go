package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=2594 lang=golang
 *
 * [2594] Minimum Time to Repair Cars
 */

// @lc code=start
func repairCars(ranks []int, cars int) int64 {
	tmp := ranks[0] * cars * cars
	result := sort.Search(tmp, func(minute int) bool {
		count := 0
		for _, rank := range ranks {
			count += int(math.Floor(math.Sqrt(float64(minute) / float64(rank))))
			if count >= cars {
				return true
			}
		}
		return false
	})

	return int64(result)
}

// @lc code=end
