package main

import "sort"

/*
 * @lc app=leetcode id=3152 lang=golang
 *
 * [3152] Special Array II
 */

// @lc code=start
func isArraySpecial(nums []int, queries [][]int) []bool {

	ranges := [][2]int{}

	tmp := [2]int{0, 0}
	for i := 1; i < len(nums); i++ {
		if nums[i]%2 != nums[i-1]%2 {
			tmp[1] = i
		} else {
			ranges = append(ranges, tmp)
			tmp[0] = i
			tmp[1] = i
		}
	}
	tmp[1] = len(nums) - 1
	ranges = append(ranges, tmp)
	result := make([]bool, len(queries))

	for i, query := range queries {

		search := sort.Search(len(ranges), func(i int) bool {
			return query[0] < ranges[i][0]
		})

		result[i] = query[1] <= ranges[search-1][1]

	}

	return result
}

// @lc code=end
