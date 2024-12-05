package main

/*
 * @lc app=leetcode id=2337 lang=golang
 *
 * [2337] Move Pieces to Obtain a String
 */

// @lc code=start
func canChange(start string, target string) bool {
	l1, r1, l2, r2 := []int{}, []int{}, []int{}, []int{}

	for i := 0; i < len(start); i++ {
		if start[i] == 'L' {
			l1 = append(l1, i)
		} else if start[i] == 'R' {
			r1 = append(r1, i)
		}
		if target[i] == 'L' {
			l2 = append(l2, i)
		} else if target[i] == 'R' {
			r2 = append(r2, i)
		}
	}
	if len(r1) != len(r2) || len(l1) != len(l2) {
		return false
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] < l2[i] {
			return false
		}
		for j := 0; j < len(r1); j++ {
			if r1[j] < l1[i] && r1[j] >= l2[i] {
				return false
			}
			if r1[j] > l2[i] {
				break
			}
		}
		l1[i] = l2[i]
	}
	for i := len(r1) - 1; i >= 0; i-- {
		if r1[i] > r2[i] {
			return false
		}
		for j := len(l1) - 1; j >= 0; j-- {
			if l1[j] > r1[i] && l1[j] <= r2[i] {
				return false
			}

			if l1[j] < r2[i] {
				break
			}
		}
		r1[i] = r2[i]
	}
	return true

}

// @lc code=end
