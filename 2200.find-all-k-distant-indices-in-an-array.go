package main

/*
 * @lc app=leetcode id=2200 lang=golang
 *
 * [2200] Find All K-Distant Indices in an Array
 */

// @lc code=start
func findKDistantIndices(nums []int, key int, k int) []int {
	res := make([]int, 0, len(nums))
	prev := -1
	for i, n := range nums {
		if n == key {
			from, to := max(i-k, 0, prev+1), min(i+k, len(nums)-1)
			for ; from <= to; from++ {
				res = append(res, from)
				prev = from
			}
		}
	}
	return res
}

// @lc code=end
