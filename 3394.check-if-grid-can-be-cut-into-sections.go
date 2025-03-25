package main

import "sort"

/*
 * @lc app=leetcode id=3394 lang=golang
 *
 * [3394] Check if Grid can be Cut into Sections
 */

// @lc code=start
func checkValidCuts(n int, rectangles [][]int) bool {
	hArr := make([][]int, len(rectangles))
	vArr := make([][]int, len(rectangles))
	for i, rectangle := range rectangles {
		hArr[i], vArr[i] = make([]int, 2), make([]int, 2)
		hArr[i][0], hArr[i][1] = rectangle[1], rectangle[3]
		vArr[i][0], vArr[i][1] = rectangle[0], rectangle[2]
	}

	var check func([][]int) bool
	check = func(arr [][]int) bool {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i][0] < arr[j][0]
		})
		count := 1
		for i := 1; i < len(arr); i++ {
			if arr[i][0] < arr[i-1][1] {
				arr[i][1] = max(arr[i][1], arr[i-1][1])
				arr[i][0] = arr[i-1][0]
			} else {
				count++
				if count == 3 {
					return true
				}
			}
		}
		return false
	}
	return check(hArr) || check(vArr)
}

// @lc code=end
