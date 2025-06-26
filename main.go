package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "AB"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))
}
