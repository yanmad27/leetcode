package main

/*
 * @lc app=leetcode id=898 lang=golang
 *
 * [898] Bitwise ORs of Subarrays
 */

// @lc code=start
func subarrayBitwiseORs(arr []int) int {
	res := make(map[int]bool)
	prevMap := make(map[int]bool)
	for _, num := range arr {
		newPrevMap := make(map[int]bool)
		newPrevMap[num] = true
		for prev := range prevMap {
			newPrevMap[num|prev] = true
		}
		prevMap = newPrevMap
		for key := range prevMap {
			res[key] = true
		}

	}
	return len(res)
}

// @lc code=end
