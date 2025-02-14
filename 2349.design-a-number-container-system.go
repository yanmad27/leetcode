package main

import "sort"

/*
 * @lc app=leetcode id=2349 lang=golang
 *
 * [2349] Design a Number Container System
 */

// @lc code=start
type NumberContainers struct {
	index    map[int]int
	group    map[int][]int
	isSorted map[int]bool
}

func Constructor1() NumberContainers {
	return NumberContainers{
		index:    make(map[int]int),
		group:    make(map[int][]int),
		isSorted: make(map[int]bool),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if prev, f := this.index[index]; f {
		if prev != number {
			this.remove(index, prev)
			this.add(index, number)
		}
	} else {
		this.add(index, number)
	}
}

func (this *NumberContainers) add(index, number int) {
	this.index[index] = number
	this.group[number] = append(this.group[number], index)
	this.isSorted[number] = false
}

func (this *NumberContainers) remove(index, prev int) {
	for i := 0; i < len(this.group[prev]); i++ {
		if this.group[prev][i] == index {
			this.group[prev] = append(this.group[prev][:i], this.group[prev][i+1:]...)
			delete(this.index, index)
			return
		}
	}
}

func (this *NumberContainers) Find(number int) int {
	if len(this.group[number]) == 0 {
		return -1
	}

	if sorted := this.isSorted[number]; !sorted {
		sort.Ints(this.group[number])
		this.isSorted[number] = true
	}

	return this.group[number][0]
}



/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
// @lc code=end
