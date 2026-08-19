package main

func main() {
	defer report()

	run(longestSubarray)

	expect(longestSubarray([]int{0, 0, 0}), 0)
	expect(longestSubarray([]int{1}), 0)
	expect(longestSubarray([]int{1, 1, 1, 1, 0}), 4)
}
