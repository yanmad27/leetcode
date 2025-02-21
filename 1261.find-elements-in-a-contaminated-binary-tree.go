package main

/*
 * @lc app=leetcode id=1261 lang=golang
 *
 * [1261] Find Elements in a Contaminated Binary Tree
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
type FindElements struct {
	memory map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	queue := []TreeNode{*root}
	memory := make(map[int]bool)
	memory[0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left != nil {
			val := cur.Val*2 + 1
			memory[val] = true
			cur.Left.Val = val
			queue = append(queue, *cur.Left)
		}
		if cur.Right != nil {
			val := cur.Val*2 + 2
			memory[val] = true
			cur.Right.Val = val
			queue = append(queue, *cur.Right)
		}
	}
	return FindElements{
		memory: memory,
	}

}

func (this *FindElements) Find(target int) bool {
	return this.memory[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
// @lc code=end
