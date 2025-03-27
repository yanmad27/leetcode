package main

/*
 * @lc app=leetcode id=2780 lang=golang
 *
 * [2780] Minimum Index of a Valid Split
 */

// @lc code=start
func minimumIndex(nums []int) int {

	count := 1
	dominant := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != dominant && count == 1 {
			dominant = nums[i]
		} else if nums[i] == dominant {
			count++
		} else {
			count--
		}
	}
	dominantIndex := make([]int, len(nums))
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == dominant {
			dominantIndex[k] = i
			k++
		}
	}
	for i := 0; i < k; i++ {
		if (i+1)*2 > dominantIndex[i]+1 &&
			(k-i-1)*2 >= len(nums)-dominantIndex[i] {
			return dominantIndex[i]
		}
	}
	return -1
}

// @lc code=end
