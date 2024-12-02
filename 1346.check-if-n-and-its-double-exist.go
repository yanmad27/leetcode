package main

/*
 * @lc app=leetcode id=1346 lang=golang
 *
 * [1346] Check If N and Its Double Exist
 */

// @lc code=start
func checkIfExist(arr []int) bool {

	doubleMap := make(map[int]bool)
	for _, num := range arr {
		if doubleMap[num*2] {
			return true
		} else if num%2 == 0 && doubleMap[num/2] {
			return true
		}
		doubleMap[num] = true
	}
	return false

}

// @lc code=end
