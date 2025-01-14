package main

/*
 * @lc app=leetcode id=2657 lang=golang
 *
 * [2657] Find the Prefix Common Array of Two Arrays
 */

// @lc code=start
func findThePrefixCommonArray(A []int, B []int) []int {
	existedA := make(map[int]bool)
	existedB := make(map[int]bool)
	counted := make(map[int]bool)
	result := make([]int, len(A))
	count := 0
	for i := 0; i < len(A); i++ {
		existedA[A[i]] = true
		existedB[B[i]] = true
		if A[i] == B[i] {
			count++
			counted[A[i]] = true
		} else {
			if existedA[B[i]] == true && !counted[B[i]] {
				count++
			}
			if existedB[A[i]] == true && !counted[A[i]] {
				count++
			}
		}
		result[i] = count

	}
	return result
}

// @lc code=end
