package main

/*
 * @lc app=leetcode id=2342 lang=golang
 *
 * [2342] Max Sum of a Pair With Equal Sum of Digits
 */

// @lc code=start

func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
func maximumSum(nums []int) int {

	n := len(nums)
	sumsMap := make(map[int][2]int)
	result := -1
	for i := 0; i < n; i++ {
		key := sumDigits(nums[i])
		old := sumsMap[key]
		if old[0] < nums[i] {
			sumsMap[key] = [2]int{nums[i], old[0]}
		} else if old[1] < nums[i] {
			sumsMap[key] = [2]int{old[0], nums[i]}
		}
		if sumsMap[key][0] > 0 && sumsMap[key][1] > 0 {
			result = max(result, sumsMap[key][0]+sumsMap[key][1])
		}
	}
	return result
}

// @lc code=end
