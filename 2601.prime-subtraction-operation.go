package main

import "sort"

/*
 * @lc app=leetcode id=2601 lang=golang
 *
 * [2601] Prime Subtraction Operation
 */

// @lc code=start
func getListPrime() []int {
	prime := []int{2}
	for i := 3; i < 1000; i++ {
		isPrime := true
		for _, p := range prime {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime = append(prime, i)
		}
	}
	return prime
}
func primeSubOperation(nums []int) bool {
	listPrime := getListPrime()
	for i, num := range nums {
		index := sort.Search(len(listPrime), func(i int) bool { return listPrime[i] >= num })
		if index > len(listPrime)-1 {
			index--
		}
		tmp := 0
		if i > 0 {
			tmp = nums[i-1]
		}
		for k := index; k >= 0; k-- {
			if num-listPrime[k] > tmp {
				nums[i] = num - listPrime[k]
				break
			}
		}
	}
	return sort.SliceIsSorted(nums, func(i, j int) bool { return nums[i] <= nums[j] })

}

// @lc code=end
