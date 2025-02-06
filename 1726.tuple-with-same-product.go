package main

/*
 * @lc app=leetcode id=1726 lang=golang
 *
 * [1726] Tuple with Same Product
 */

// @lc code=start
func tupleSameProduct(nums []int) int {

	count := 0
	n := len(nums)
	multiMap := make(map[int]int)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			multiMap[nums[i]*nums[j]]++
		}
	}
	for _, v := range multiMap {
		if v > 1 {
			count += v * (v - 1) * 4
		}
	}

	return count

}

// @lc code=end
