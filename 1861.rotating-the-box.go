package main

/*
 * @lc app=leetcode id=1861 lang=golang
 *
 * [1861] Rotating the Box
 */

// @lc code=start
func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])

	res := make([][]byte, n)
	for i := range res {
		res[i] = make([]byte, m)
		for j := range res[i] {
			res[i][j] = '.'
		}
	}

	for r := 0; r < m; r++ {
		i := n - 1
		for c := n - 1; c >= 0; c-- {
			if box[r][c] == '#' {
				res[i][m-r-1] = '#'
				i--
			} else if box[r][c] == '*' {
				res[c][m-r-1] = '*'
				i = c - 1
			}
		}
	}
	return res

}

// @lc code=end
