package leetcode

/*
 * @lc app=leetcode id=2326 lang=golang
 *
 * [2326] Spiral Matrix IV
 */

// @lc code=start
func spiralMatrix(m int, n int, head *ListNode) [][]int {

	res := [][]int{}
	for i := 0; i < m; i++ {
		tmp := []int{}
		for j := 0; j < n; j++ {
			tmp = append(tmp, -1)
		}
		res = append(res, tmp)
	}

	t0, t1, t2 := 1, 0, 1
	i, j := 0, 0
	dir := 0
	for head != nil {
		res[i][j] = head.Val
		i += t0 * t1
		j += t0 * t2
		if i >= m || i < 0 || j >= n || j < 0 || res[i][j] != -1 {
			t1, t2 = t2, t1
			if dir == 0 {
				j--
				i++
				dir = 1
			} else if dir == 1 {
				i--
				j--
				t0 *= -1
				dir = 2
			} else if dir == 2 {
				j++
				i--
				dir = 3
			} else {
				i++
				j++
				t0 *= -1
				dir = 0
			}
		}
		head = head.Next
	}

	return res

}

// @lc code=end
