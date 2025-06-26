package main

/*
 * @lc app=leetcode id=2200 lang=golang
 *
 * [2200] Find All K-Distant Indices in an Array
 */

// @lc code=start
func findKDistantIndices(nums []int, key int, k int) []int {
	rs := make([]int, 0, len(nums))
	pre := -1
	for i, n := range nums {
		if n == key {
			j, k := max(i-k, 0, pre+1), min(i+k, len(nums)-1)
			for ; j <= k; j++ {
				rs = append(rs, j)
				pre = j
			}
		}
	}
	return rs
}

// @lc code=end
