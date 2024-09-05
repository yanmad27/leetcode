package main

/*
 * @lc app=leetcode id=2028 lang=golang
 *
 * [2028] Find Missing Observations
 */

// @lc code=start
func missingRolls(rolls []int, mean int, n int) []int {
	sum := mean * (len(rolls) + n)
	for _, roll := range rolls {
		sum -= roll
	}
	var find func(tmp []int) []int
	find = func(tmp []int) []int {
		if len(tmp) == n {
			tmpSum := 0
			for _, num := range tmp {
				tmpSum += num
			}
			if tmpSum == sum {
				return tmp
			}
			return nil
		}
		for i := sum / n; i <= 6; i++ {
			result := find(append(tmp, i))
			if len(result) != 0 {
				return result
			}
		}
		return []int{}
	}
	return find([]int{})
}

// @lc code=end
