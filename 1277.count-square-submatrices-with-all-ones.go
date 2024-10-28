package main

/*
 * @lc app=leetcode id=1277 lang=golang
 *
 * [1277] Count Square Submatrices with All Ones
 */

// @lc code=start

func countSquaresByLevel(matrix [][]int, level int) int {
	count := 0
	m, n := len(matrix), len(matrix[1])
	for i := 0; i < m-level+1; i++ {
		for j := 0; j < n-level+1; j++ {
			isValid := true
			for k := 0; k < level; k++ {
				for h := 0; h < level; h++ {
					if matrix[i+k][j+h] == 0 {
						isValid = false
						break
					}
				}
				if isValid == false {
					break
				}
			}

			if isValid {
				count++
			}
		}
	}

	return count
}
func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[1])

	count := 0
	maxSize := min(m, n) + 1
	for level := 1; level < maxSize; level++ {
		count += countSquaresByLevel(matrix, level)
	}
	return count
}

// @lc code=end
