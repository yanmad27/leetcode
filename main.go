package main

import "fmt"

func main() {
	fmt.Println(getWordsInLongestSubsequence([]string{"cdb", "cdd", "cd", "dcc", "cca", "cda", "ca", "cc", "bcc"}, []int{8, 5, 9, 5, 2, 7, 4, 7, 3}))
	// fmt.Println(getWordsInLongestSubsequence([]string{"bab", "dab", "cab"}, []int{1, 2, 2}))
}
