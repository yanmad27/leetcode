package main

/*
 * @lc app=leetcode id=2471 lang=golang
 *
 * [2471] Minimum Number of Operations to Sort a Binary Tree by Level
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMinOperator(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
			count++
		}
	}
	return count
}
func minimumOperations(root *TreeNode) int {
	valueByLevel := make(map[int][]int)
	var travel func(root *TreeNode, level int)
	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		valueByLevel[level] = append(valueByLevel[level], root.Val)
		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}
	travel(root, 0)
	count := 0
	for _, value := range valueByLevel {
		count += findMinOperator(value)
	}
	return count
}

// @lc code=end
