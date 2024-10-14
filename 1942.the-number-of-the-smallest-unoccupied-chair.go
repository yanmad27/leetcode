package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=1942 lang=golang
 *
 * [1942] The Number of the Smallest Unoccupied Chair
 */

// @lc code=start

type priorityQueue []int

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }

func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x int) { *pq = append(*pq, x) }

func (pq *priorityQueue) Pop() int {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq *priorityQueue) peek() int { return (*pq)[0] }

func smallestUnoccupiedChair(n int, times [][]int) int {
	chairs := make([]int, n)
	for i := range chairs {
		chairs[i] = i + 1
	}
	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})
	available := make(priorityQueue, 0, n)
	occupied := make(map[int]int)
	for _, time := range times {
		arrival, leave := time[0], time[1]
		for available.Len() > 0 && occupied[available.peek()] <= arrival {
			chair := available.Pop()
			delete(occupied, chair)
		}
		for len(chairs) > 0 {
			chair := chairs[0]
			chairs = chairs[1:]
			occupied[chair] = leave
			available.Push(chair)
			break
		}
	}
	return available.peek()
}

// @lc code=end
