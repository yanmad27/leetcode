package main

/*
 * @lc app=leetcode id=2364 lang=golang
 *
 * [2364] Count Number of Bad Pairs
 */

// @lc code=start
func countBadPairs(nums []int) int64 {
	n := len(nums)
	hashMap := make(map[int]int)
	count := 0
	for i, num := range nums {
		count += hashMap[num-i]
		hashMap[num-i]++
	}
	return int64(n*(n-1)/2 - count)

}

// @lc code=end
