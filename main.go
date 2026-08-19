package main

func main() {
	defer report()
	expect(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4), 12.75)
	expect(findMaxAverage([]int{5}, 1), 5)
}
