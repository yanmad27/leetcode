package main

/*
* @lc app=leetcode id=1590 lang=golang
 *
 * [1590] Make Sum Divisible by P
*/

// @lc code=start
func minSubarray(nums []int, p int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	rest := sum % p
	if rest == 0 {
		return 0
	}
	if sum < p {
		return -1
	}
	mod := make([]int, len(nums))
	for i, num := range nums {
		mod[i] = num % p
	}
	l, r := 0, 0
	tmp := 0
	nums = append(nums, 0)
	minSub := len(nums)
	for r < len(nums) {
		if tmp%p == rest {
			minSub = min(r-l, minSub)
			tmp -= mod[l]
			l++
		} else if tmp < p {
			tmp += mod[r]
			r++
		} else if tmp > p {
			tmp -= mod[l]
			l++
		}
	}
	if minSub == len(nums) {
		return -1
	}
	return minSub
}

// @lc code=end
