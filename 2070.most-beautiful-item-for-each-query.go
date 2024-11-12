package main

import "sort"

/*
 * @lc app=leetcode id=2070 lang=golang
 *
 * [2070] Most Beautiful Item for Each Query
 */

// @lc code=start
func maximumBeauty(items [][]int, queries []int) []int {

	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] < items[j][1]
		}
		return items[i][0] < items[j][0]
	})
	prices := []int{items[0][0]}
	beauty := []int{items[0][1]}
	maxBeauty := items[0][1]
	for i := 1; i < len(items); i++ {
		if items[i][1] > maxBeauty {
			maxBeauty = items[i][1]
		}
		if items[i][0] != items[i-1][0] {
			prices = append(prices, items[i][0])
			beauty = append(beauty, maxBeauty)
		} else {
			prices[len(prices)-1] = items[i][0]
			beauty[len(beauty)-1] = maxBeauty
		}
	}
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		index := sort.SearchInts(prices, queries[i])
		if index == len(prices) || queries[i] < prices[index] {
			index--
		}
		if index == -1 {
			ans[i] = 0
		} else {
			ans[i] = beauty[index]
		}
	}
	return ans
}

// @lc code=end
