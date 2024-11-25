package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=773 lang=golang
 *
 * [773] Sliding Puzzle
 */

// @lc code=start
func findShortestPath(start, target string) int {
	if target == start {
		return 0
	}
	moves := [6][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}
	queue := make([]string, 0)
	count := 0
	queue = append(queue, start)
	visited := map[string]bool{start: true}
	for len(queue) > 0 {
		length := len(queue)

		for length > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return count
			}
			zeroIndex := strings.IndexRune(cur, '0')
			for _, move := range moves[zeroIndex] {
				state := []byte(cur)
				state[zeroIndex], state[move] = state[move], state[zeroIndex]
				nextState := string(state)
				if !visited[nextState] {
					visited[nextState] = true
					queue = append(queue, nextState)
				}
			}
			length--
		}
		count++
	}
	return -1
}

func slidingPuzzle(board [][]int) int {
	target := "123450"
	start := ""

	for _, row := range board {
		for _, val := range row {
			start = start + fmt.Sprintf("%d", val)
		}
	}
	return findShortestPath(start, target)
}

// @lc code=end
