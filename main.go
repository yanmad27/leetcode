package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome([]string{"lc", "cl", "gg"}))
	fmt.Println(longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
	fmt.Println(longestPalindrome([]string{"cc", "ll", "xx"}))
	fmt.Println(longestPalindrome([]string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}))
}
