package main

import "fmt"

/*
 * @lc app=leetcode id=3043 lang=golang
 *
 * [3043] Find the Length of the Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixMap := make(map[int]bool)
	for _, num := range arr1 {
		for num != 0 {
			prefixMap[num] = true
			num /= 10
		}
	}
	result := 0

	for _, num := range arr2 {
		for num != 0 {
			if prefixMap[num] {
				result = max(result, len(fmt.Sprintf("%d", num)))
				break
			}
			num /= 10
		}
	}
	return result
}

// @lc code=end
