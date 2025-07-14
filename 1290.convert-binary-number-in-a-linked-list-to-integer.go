package main

import "math"

/*
 * @lc app=leetcode id=1290 lang=golang
 *
 * [1290] Convert Binary Number in a Linked List to Integer
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	n := len(result)
	base10 := 0
	for i := 0; i < n; i++ {
		base10 += result[i] * int(math.Pow(float64(2), float64(n-i-1)))
	}
	return base10
}

// @lc code=end
