package main

/*
 * @lc app=leetcode id=2206 lang=golang
 *
 * [2206] Divide Array Into Equal Pairs
 */

// @lc code=start
func divideArray(nums []int) bool {
	n := len(nums)
	numMap := make(map[int]int, n)
	for _, num := range nums {
		numMap[num]++
	}
	for _, val := range numMap {
		if val%2 == 1 {
			return false
		}
	}
	return true
}

// @lc code=end
