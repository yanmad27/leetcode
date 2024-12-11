package main

/*
 * @lc app=leetcode id=2337 lang=golang
 *
 * [2337] Move Pieces to Obtain a String
 */

// @lc code=start
func canChange(start string, target string) bool {
	t1, t2 := []byte{}, []byte{}
	l1, r1, l2, r2 := []int{}, []int{}, []int{}, []int{}

	for i := 0; i < len(start); i++ {
		if start[i] == 'L' {
			l1 = append(l1, i)
			t1 = append(t1, 'L')
		} else if start[i] == 'R' {
			r1 = append(r1, i)
			t1 = append(t1, 'R')
		}
		if target[i] == 'L' {
			l2 = append(l2, i)
			t2 = append(t2, 'L')
		} else if target[i] == 'R' {
			r2 = append(r2, i)
			t2 = append(t2, 'R')
		}
	}
	if len(r1) != len(r2) ||
		len(l1) != len(l2) ||
		string(t1) != string(t2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] < l2[i] {
			return false
		}
	}
	for i := len(r1) - 1; i >= 0; i-- {
		if r1[i] > r2[i] {
			return false
		}
	}
	return true

}

// @lc code=end
