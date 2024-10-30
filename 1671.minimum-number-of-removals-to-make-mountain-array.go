package main

import (
	"slices"
	"sort"
)

/*
 * @lc app=leetcode id=1671 lang=golang
 *
 * [1671] Minimum Number of Removals to Make Mountain Array
 */

// @lc code=start
func lengthOfLISWithMax(nums []int, max int) int {
	tmp := []int{}
	for _, num := range nums {
		if num >= max {
			continue
		}
		index := sort.SearchInts(tmp, num)
		if index == len(tmp) {
			tmp = append(tmp, num)
		} else {
			tmp[index] = num
		}
	}
	return len(tmp)
}

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	maxLength := 0
	for i := 1; i < n-1; i++ {
		cur := nums[i]
		left := nums[:i]
		right := nums[i+1:]
		slices.Reverse(right)
		lengthLeft := lengthOfLISWithMax(left, cur)
		lengthRight := lengthOfLISWithMax(right, cur)
		if lengthLeft != 0 && lengthRight != 0 {
			mountainLength := lengthLeft + lengthRight + 1
			maxLength = max(mountainLength, maxLength)
		}
	}
	return n - maxLength
}

// @lc code=end
