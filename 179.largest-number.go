package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start
func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = fmt.Sprintf("%d", num)
	}
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})
	result := strings.Join(strNums, "")
	if result[0] == '0' {
		return "0"
	}
	return result
}

// @lc code=end
