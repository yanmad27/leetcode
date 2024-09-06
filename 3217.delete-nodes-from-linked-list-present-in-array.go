package leetcode

/*
@lc app=leetcode id=3217 lang=golang
 *
 * [3217] Delete Nodes From Linked List Present in Array
*/

// @lc code=start

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}
	for numsMap[head.Val] {
		head = head.Next
	}
	pre := head
	cur := head.Next
	for cur != nil {
		if _, ok := numsMap[cur.Val]; ok {
			cur = cur.Next
			pre.Next = cur
		} else {
			cur = cur.Next
			pre = pre.Next
		}
	}
	return head

}

// @lc code=end
