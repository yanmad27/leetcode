package problems

import (
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

// run calls fn against every example in the problem description, so a test is
// just:
//
//	run(longestSubarray)
//
// The problem id comes from the "@lc ... id=" header of the file where fn is
// defined, and the description comes from the cache the LeetCode extension
// writes under ~/.lc when you open the problem.
//
// Cases of your own are plain calls, type-checked like any other Go code:
//
//	expect(longestSubarray([]int{0, 0, 0}), 0)
func run(fn any) {
	desc, err := describe(fn)
	if err != nil {
		panic(fmt.Sprintf("run: %v", err))
	}
	runText(fn, desc)
}

// runText is the fallback for a problem run cannot find a description for:
// paste the examples verbatim from LeetCode.
//
//	runText(longestSubarray, `
//	Input: nums = [1,1,0,1]
//	Output: 3
//	Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
//	`)
//
// Each "Input:" line is paired with the "Output:" line that follows it; blank
// lines and Explanation lines are ignored. Values are decoded as JSON into the
// function's parameter and return types, so numbers, strings, bools and nested
// slices all work. Examples that do not fit that shape (design problems, tree
// and linked-list inputs) are reported and skipped.
func runText(fn any, examples string) {
	fv := reflect.ValueOf(fn)
	ft := fv.Type()
	if ft.Kind() != reflect.Func || ft.NumOut() != 1 {
		panic("run: want a function with exactly one return value")
	}

	cases := parseCases(examples, ft)
	width := 0
	for _, c := range cases {
		if n := len(c.input); n > width && n <= inputColumn {
			width = n
		}
	}
	for _, c := range cases {
		expectInput(fv.Call(c.args)[0].Interface(), c.want, c.input, width)
	}
}

// inputColumn caps how far the expected/got halves are pushed right. One long
// example does not get to indent every other line past it.
const inputColumn = 44

type example struct {
	input string
	args  []reflect.Value
	want  any
}

// parseCases reads the "Input:"/"Output:" pairs out of a problem description.
// Examples that do not fit the shape of fn are reported and dropped, so that
// the cases returned are all runnable.
func parseCases(examples string, ft reflect.Type) []example {
	var cases []example
	var args []reflect.Value
	var input string
	ok := true

	for _, line := range strings.Split(examples, "\n") {
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "Input:"):
			input = strings.TrimSpace(strings.TrimPrefix(line, "Input:"))
			args, ok = parseArgs(input, ft)
		case strings.HasPrefix(line, "Output:") && ok:
			want, err := decode(strings.TrimSpace(strings.TrimPrefix(line, "Output:")), ft.Out(0))
			if err != nil {
				fmt.Printf("SKIP: %s: %v\n", input, err)
				continue
			}
			cases = append(cases, example{input, args, want})
		}
	}
	return cases
}

func parseArgs(input string, ft reflect.Type) ([]reflect.Value, bool) {
	fields := splitTopLevel(input)
	if len(fields) != ft.NumIn() {
		fmt.Printf("SKIP: %s: %s takes %d args, example has %d\n",
			input, ft, ft.NumIn(), len(fields))
		return nil, false
	}
	args := make([]reflect.Value, len(fields))
	for i, field := range fields {
		v, err := decode(stripName(field), ft.In(i))
		if err != nil {
			fmt.Printf("SKIP: %s: %v\n", input, err)
			return nil, false
		}
		args[i] = reflect.ValueOf(v)
	}
	return args, true
}

// stripName drops the "nums = " part of an argument, keeping only the value.
func stripName(field string) string {
	if i := strings.Index(field, "="); i >= 0 && !strings.ContainsAny(field[:i], `["{`) {
		return strings.TrimSpace(field[i+1:])
	}
	return field
}

func decode(s string, t reflect.Type) (any, error) {
	v := reflect.New(t)
	if err := json.Unmarshal([]byte(s), v.Interface()); err != nil {
		return nil, fmt.Errorf("cannot read %q as %s", s, t)
	}
	return v.Elem().Interface(), nil
}

// splitTopLevel splits on commas that sit outside brackets and quotes, so
// `nums = [1,2], k = 3` yields two fields rather than three.
func splitTopLevel(s string) []string {
	var fields []string
	depth, quoted, start := 0, false, 0
	for i, r := range s {
		switch {
		case quoted:
			if r == '"' {
				quoted = false
			}
		case r == '"':
			quoted = true
		case r == '[' || r == '{':
			depth++
		case r == ']' || r == '}':
			depth--
		case r == ',' && depth == 0:
			fields = append(fields, strings.TrimSpace(s[start:i]))
			start = i + 1
		}
	}
	if rest := strings.TrimSpace(s[start:]); rest != "" {
		fields = append(fields, rest)
	}
	return fields
}

var (
	idHeader = regexp.MustCompile(`@lc .*id=(\d+)`)
	idPrefix = regexp.MustCompile(`^(\d+)\.`)
	htmlTag  = regexp.MustCompile(`<[^>]*>`)
)

// describe returns the problem statement as plain text, read from the cache the
// LeetCode extension keeps for the problem that fn solves.
func describe(fn any) (string, error) {
	pc := reflect.ValueOf(fn).Pointer()
	f := runtime.FuncForPC(pc)
	if f == nil {
		return "", fmt.Errorf("cannot locate the source file of %T", fn)
	}
	src, _ := f.FileLine(pc)

	id, err := problemID(src)
	if err != nil {
		return "", err
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	for _, site := range []string{"leetcode", "leetcode.cn"} {
		matches, _ := filepath.Glob(filepath.Join(home, ".lc", site, "cache", id+".*.json"))
		for _, match := range matches {
			raw, err := os.ReadFile(match)
			if err != nil {
				continue
			}
			var problem struct{ Desc string }
			if err := json.Unmarshal(raw, &problem); err != nil {
				continue
			}
			return html.UnescapeString(htmlTag.ReplaceAllString(problem.Desc, "")), nil
		}
	}
	return "", fmt.Errorf("no cached description for problem %s; open it in the LeetCode extension first", id)
}

// problemID reads the id from the "@lc ... id=1493" header the extension writes,
// falling back to the "1493." prefix of the file name.
func problemID(src string) (string, error) {
	if raw, err := os.ReadFile(src); err == nil {
		if m := idHeader.FindSubmatch(raw); m != nil {
			return string(m[1]), nil
		}
	}
	if m := idPrefix.FindStringSubmatch(filepath.Base(src)); m != nil {
		return m[1], nil
	}
	return "", fmt.Errorf("cannot find a problem id in %s", filepath.Base(src))
}
