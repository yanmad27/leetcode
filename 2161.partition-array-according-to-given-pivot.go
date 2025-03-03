package main

/*
 * @lc app=leetcode id=2161 lang=golang
 *
 * [2161] Partition Array According to Given Pivot
 */

// @lc code=start
func pivotArray(nums []int, pivot int) []int {

	les := make([]int, 0, len(nums))
	mid := make([]int, 0, len(nums))
	lar := make([]int, 0, len(nums))

	for _, num := range nums {
		if num < pivot {
			les = append(les, num)
		} else if num == pivot {
			mid = append(mid, num)
		} else {
			lar = append(lar, num)
		}
	}
	les = append(les, mid...)
	les = append(les, lar...)
	return les
}

// @lc code=end
