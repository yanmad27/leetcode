# leetcode

Go solutions, one file per problem under `problems/`. Test cases come from the
problem description itself, so there is nothing to write by hand.

## Running

```sh
./solve       # the solution file you edited most recently
./solve 724   # problem 724
```

```
$ ./solve 1493
run(longestSubarray)   [problems/1493.longest-subarray-of-1-s-after-deleting-one-element.go]
3/3 passed
```

A failing case prints what went in and what came back:

```
$ ./solve 724
run(pivotIndex)   [problems/724.find-pivot-index.go]
FAIL #1: expected 3 (int), but got 0 (int)
  input: nums = [1,7,3,6,5,6]
1/3 passed
```

## How it works

`run(fn)` finds the problem id in the `@lc app=leetcode id=724` header of the
file where `fn` is defined, reads the description the LeetCode VS Code extension
cached under `~/.lc/leetcode/cache/`, and runs every `Input:`/`Output:` pair in
it. Values are decoded as JSON into the function's own parameter and return
types, so ints, floats, strings, bools and nested slices all work. Numbers
compare across types, so an `int` expectation matches a `float64` result.

Examples that do not fit that shape — trees, linked lists, `[]byte`, design
problems — are reported and skipped rather than crashing the run:

```
SKIP: cannot read "[5,4,8,11,null,13,4,7,2,null,null,null,1]" as *main.TreeNode
```

Two things follow from Go not being able to look a function up by name at run
time:

- `./solve` has to compile the call in. It writes `solve_current.go`, runs, and
  deletes it, so nothing new is left in the repo. `main.go` never changes.
- Two problems that share a function name will not compile together. Rename one
  (`longestSubarray1493`); `run` keys off the file, not the name, so it keeps
  working.

If a problem is not in the cache, open it once in the extension.

## Cases of your own

The description rarely covers the edge cases that actually break a solution.
Add a `Cases` function next to the solution, named after it, and `./solve` calls
it after the examples:

```go
// in problems/1493.longest-subarray-of-1-s-after-deleting-one-element.go
func longestSubarrayCases() {
	expect(longestSubarray([]int{0, 0, 0}), 0)
	expect(longestSubarray([]int{1}), 0)
}
```

```
$ ./solve 1493
run(longestSubarray)   [problems/1493.longest-subarray-of-1-s-after-deleting-one-element.go]
5/5 passed
```

Plain Go, so the compiler checks the types, and the cases are committed with the
solution they belong to. Problems without such a function just run the examples.

## Layout

Everything lives in package `problems`, so a solution can call `expect` and the
shared types without importing anything. `main.go` only starts it.

| | |
|---|---|
| `solve` | picks a problem and runs it |
| `main.go` | entry point, calls `problems.Main` |
| `problems/724.find-pivot-index.go` | one problem, written by the LeetCode extension |
| `problems/run.go` | reads examples out of the cached description |
| `problems/expect.go` | comparison and the pass/total tally |
| `problems/entry.go` | holds the problem `./solve` selected |
| `problems/model.go`, `problems/heap.go` | shared `TreeNode`, `ListNode`, heap helpers |

Point the LeetCode extension at the folder so new files land there:
`"leetcode.filePath": { "default": { "folder": "problems" } }`, and write
`package problems` at the top.
