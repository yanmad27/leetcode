package main

/*
 * @lc app=leetcode id=1652 lang=golang
 *
 * [1652] Defuse the Bomb
 */

// @lc code=start

func decrypt(code []int, k int) []int {

	n := len(code)
	tmp := []int{}
	tmp = append(tmp, code...)
	tmp = append(tmp, code...)
	rs := make([]int, n)
	if k > 0 {
		sum := 0
		for i := 1; i <= k; i++ {
			sum += code[i]
		}
		for i := 0; i < n; i++ {
			rs[i] = sum
			sum -= tmp[i+1]
			sum += tmp[i+k+1]
		}
	} else if k < 0 {
		sum := 0
		for i := n - 2; i >= n-2+k+1; i-- {
			sum += code[i]
		}
		for i := n - 1; i >= 0; i-- {
			rs[i] = sum
			sum -= tmp[n+i-1]
			sum += tmp[n+i+k-1]
		}
	}
	return rs
}

// @lc code=end
