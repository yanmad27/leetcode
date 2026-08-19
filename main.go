package main

func main() {
	defer report()
	expect(maxOperations([]int{1, 2, 3, 4}, 5), 1)
	expect(maxOperations([]int{3, 1, 3, 4, 3}, 3), 1)
}
