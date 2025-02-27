package main

/*
 * @lc app=leetcode id=873 lang=golang
 *
 * [873] Length of Longest Fibonacci Subsequence
 */

// @lc code=start
func lenLongestFibSubseq(arr []int) int {

	n := len(arr)
	result := 0
	mem := map[int]bool{}
	for _, num := range arr {
		mem[num] = true
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {

			length := 0
			a, b := arr[i], arr[j]
			for mem[a+b] {
				a, b = b, a+b
				length++
			}
			if length > 0 {
				result = max(result, length+2)
			}
		}
	}
	return result
}

// @lc code=end
