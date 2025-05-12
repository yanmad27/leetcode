package main

/*
 * @lc app=leetcode id=2094 lang=golang
 *
 * [2094] Finding 3-Digit Even Numbers
 */

// @lc code=start
func findEvenNumbers(digits []int) []int {
	countMap := make([]int, 10)
	for _, digit := range digits {
		countMap[digit]++
	}

	var canForm func(int) bool
	canForm = func(num int) bool {
		countMap2 := make([]int, 10)
		for num > 0 {
			digit := num % 10
			countMap2[digit]++
			if countMap2[digit] > countMap[digit] {
				return false
			}
			num /= 10
		}
		return true
	}
	result := []int{}
	for i := 100; i < 1000; i += 2 {
		if canForm(i) {
			result = append(result, i)
		}
	}
	return result
}

// @lc code=end
