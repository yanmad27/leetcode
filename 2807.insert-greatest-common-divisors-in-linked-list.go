package leetcode

/*
 * @lc app=leetcode id=2807 lang=golang
 *
 * [2807] Insert Greatest Common Divisors in Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	result := head
	for head.Next != nil {
		middle := &ListNode{
			Val:  findGCD(head.Val, head.Next.Val),
			Next: head.Next,
		}
		head.Next = middle
		head = middle.Next
	}
	return result
}

// @lc code=end
