package main

/*
 * @lc app=leetcode id=1760 lang=golang
 *
 * [1760] Minimum Limit of Balls in a Bag
 */

// @lc code=start

func minimumSize(nums []int, maxOps int) int {
	low, high := 1, 0
	for _, n := range nums {
		if n > high {
			high = n
		}
	}

	for low < high {
		mid := (low + high) / 2
		ops := 0
		for _, n := range nums {
			ops += (n - 1) / mid
		}
		if ops <= maxOps {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}

// @lc code=end
