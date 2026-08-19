package problems

/*
 * @lc app=leetcode id=1386 lang=golang
 *
 * [1386] Cinema Seat Allocation
 */

// @lc code=start

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// A family of four fits in one of three blocks, and seats 1 and 10 are too
	// close to the walls to be part of any of them:
	//
	//	1 [2 3 4 5] 6 7 8 9 10
	//	1 2 3 [4 5 6 7] 8 9 10
	//	1 2 3 4 5 [6 7 8 9] 10
	//
	// left and right do not overlap, so a row seats two families when both are
	// clear, and one when any single block is.
	const (
		left   = 1<<2 | 1<<3 | 1<<4 | 1<<5
		middle = 1<<4 | 1<<5 | 1<<6 | 1<<7
		right  = 1<<6 | 1<<7 | 1<<8 | 1<<9
	)

	// n reaches 1e9, so the rows are never walked. Only the ones that hold a
	// reservation that matters are collected, each as a bitmask of its seats.
	taken := make(map[int]int)
	for _, seat := range reservedSeats {
		row, col := seat[0], seat[1]
		if col >= 2 && col <= 9 {
			taken[row] |= 1 << col
		}
	}

	// Every row left untouched seats two families; the rest are counted by hand.
	families := 2 * (n - len(taken))
	for _, seats := range taken {
		switch {
		case seats&left == 0 && seats&right == 0:
			families += 2
		case seats&left == 0 || seats&middle == 0 || seats&right == 0:
			families++
		}
	}
	return families
}

// @lc code=end
