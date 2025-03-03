package main
/*
 * @lc app=leetcode id=2161 lang=golang
 *
 * [2161] Partition Array According to Given Pivot
 */

// @lc code=start
func pivotArray(nums []int, pivot int) []int {
    
	Lin
	les := []int{}
	mid:=[]int{}
	lar:= []int{}

	for _, num:= range nums {
		if num < pivot {
			les = append(les, num)
		} else if num == pivot {
			mid = append(mid, num)
		} else {
			lar = append(lar, num)
		}
	}
	rs := []int{}
	rs = append(rs, les...)
	rs = append(rs, mid...)
	rs = append(rs, lar...)
	return  rs
}
// @lc code=end

