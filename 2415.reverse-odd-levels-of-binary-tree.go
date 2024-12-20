package main

import (
	"math"
)

/*
 * @lc app=leetcode id=2415 lang=golang
 *
 * [2415] Reverse Odd Levels of Binary Tree
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
func reverseOddLevels(root *TreeNode) *TreeNode {

	arr := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		arr = append(arr, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	maxLevel := int(math.Ceil(math.Sqrt(float64(len(arr)))))
	for level := 1; level <= maxLevel; level += 2 {
		start := int(math.Pow(2, float64(level))) - 1
		end := int(math.Pow(2, float64(level+1))) - 2
		for start < end && end < len(arr) {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}
	queue = []*TreeNode{root}
	idx := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Val = arr[idx]
		idx++
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}

// @lc code=end
