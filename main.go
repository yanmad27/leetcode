package main

import "fmt"

func main() {
	rolls := []int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}
	mean := 4
	n := 40
	fmt.Println(missingRolls(rolls, mean, n))
}
