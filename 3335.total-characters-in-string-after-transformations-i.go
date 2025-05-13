package main

/*
 * @lc app=leetcode id=3335 lang=golang
 *
 * [3335] Total Characters in String After Transformations I
 */

// @lc code=start
const mod = 1000000007

func lengthAfterTransformations(s string, t int) int {
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	for round := 1; round <= t; round++ {
		next := make([]int, 26)
		for c := 0; c < 25; c++ {
			next[c+1] += count[c]
		}
		next[0] = count[25] % mod
		next[1] += count[25] % mod
		count = next
	}
	result := 0
	for i := 0; i < 26; i++ {
		result += count[i]
	}
	return result % mod
}

// @lc code=end
