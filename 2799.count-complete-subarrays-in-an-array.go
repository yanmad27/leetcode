package main

/*
 * @lc app=leetcode id=2799 lang=golang
 *
 * [2799] Count Complete Subarrays in an Array
 */

// @lc code=start
func countCompleteSubarrays(nums []int) int {
	distinct := make(map[int]bool)
	for _, num := range nums {
		distinct[num] = true
	}
	n := len(nums)
	nd := len(distinct)
	count := 0
	for l := nd; l <= n; l++ {
		miniDistinct := make(map[int]int, l)
		for i := range l {
			miniDistinct[nums[i]]++
		}
		if len(miniDistinct) == nd {
			count++
		}

		for i := l; i < n; i++ {
			miniDistinct[nums[i-l]]--
			if miniDistinct[nums[i-l]] == 0 {
				delete(miniDistinct, nums[i-l])
			}
			miniDistinct[nums[i]]++
			if len(miniDistinct) == nd {
				count++
			}
		}

	}
	return count
}

// @lc code=end
