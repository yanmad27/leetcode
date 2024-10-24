package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=2583 lang=golang
 *
 * [2583] Kth Largest Sum in a Binary Tree
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

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	sumMap := make(map[int]int)
	var recursive func(root *TreeNode, level int)
	recursive = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		sumMap[level] += root.Val
		recursive(root.Left, level+1)
		recursive(root.Right, level+1)
	}
	recursive(root, 0)
	if k > len(sumMap) {
		return -1
	}
	sumArr := []int{}
	for _, sum := range sumMap {
		sumArr = append(sumArr, sum)
	}
	sort.Ints(sumArr)
	return int64(sumArr[len(sumArr)-k])
}

// @lc code=end
