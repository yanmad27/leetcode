package problems

/*
 * @lc app=leetcode id=2215 lang=golang
 *
 * [2215] Find the Difference of Two Arrays
 */

// @lc code=start
func findDifference(nums1 []int, nums2 []int) [][]int {
	// The sets answer "is this value in the other array?" in O(1).
	in1 := make(map[int]bool, len(nums1))
	for _, n := range nums1 {
		in1[n] = true
	}
	in2 := make(map[int]bool, len(nums2))
	for _, n := range nums2 {
		in2[n] = true
	}

	// Empty rather than nil, so a side with nothing in it encodes as [], not null.
	only1, only2 := []int{}, []int{}

	// Both halves walk the original slice rather than the set: Go randomises map
	// iteration, so ranging over in1 would shuffle the answer between runs. taken
	// carries the de-duplication that ranging over a set would have given.
	taken := make(map[int]bool)
	for _, n := range nums1 {
		if !in2[n] && !taken[n] {
			taken[n] = true
			only1 = append(only1, n)
		}
	}
	clear(taken)
	for _, n := range nums2 {
		if !in1[n] && !taken[n] {
			taken[n] = true
			only2 = append(only2, n)
		}
	}

	return [][]int{only1, only2}
}

// @lc code=end
