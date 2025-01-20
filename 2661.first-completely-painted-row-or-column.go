package main

/*
 * @lc app=leetcode id=2661 lang=golang
 *
 * [2661] First Completely Painted Row or Column
 */

// @lc code=start
func firstCompleteIndex(arr []int, mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	rows := make([][]int, m)
	cols := make([][]int, n)
	indexMap := make(map[int][2]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			indexMap[mat[i][j]] = [2]int{i, j}
		}
	}
	for rs, num := range arr {
		index := indexMap[num]
		i, j := index[0], index[1]
		rows[i] = append(rows[i], j)
		cols[j] = append(cols[j], i)
		if len(rows[i]) == n {
			return rs
		}
		if len(cols[j]) == m {
			return rs
		}
	}
	panic("unreachable")
}

// @lc code=end
