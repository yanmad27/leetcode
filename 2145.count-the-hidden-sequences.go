package main

/*
 * @lc app=leetcode id=2145 lang=golang
 *
 * [2145] Count the Hidden Sequences
 */

// @lc code=start
func numberOfArrays(differences []int, lower int, upper int) int {
	maxDiff := differences[0]
	minDiff := differences[0]
	cur := 0
	for _, num := range differences {
		cur += num
		maxDiff = max(maxDiff, cur)
		minDiff = min(minDiff, cur)
	}

	count := 0
	for i := lower; i <= upper; i++ {
		if i+maxDiff <= upper && i+minDiff >= lower {
			count++
		}
	}
	return count
}

// @lc code=end
