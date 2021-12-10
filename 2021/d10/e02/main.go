package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"sort"
	"strings"
)

var openChars = "([{<"
var closeChars = ")]}>"

var scores = map[byte]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func repairScore(stack []byte) int {
	res := 0
	for i := len(stack) - 1; i >= 0; i-- {
		res *= 5
		res += scores[stack[i]]
	}
	return res
}

func scoreForLine(line string) int {
	var stack []byte
	var head byte
	for _, c := range []byte(line) {
		if idx := strings.IndexByte(closeChars, c); idx != -1 {
			oChar := openChars[idx]
			n := len(stack) - 1
			head, stack = stack[n], stack[:n]
			if oChar != head {
				return 0
			}
		} else {
			stack = append(stack, c)
		}
	}
	// Doesn't have syntax error
	return repairScore(stack)
}

func solve(input []string) int {
	var scores []int
	for _, l := range input {
		if s := scoreForLine(l); s != 0 {
			scores = append(scores, s)
		}
	}
	sort.Sort(sort.IntSlice(scores))
	return scores[len(scores)/2]
}

func main() {
	input := files.ReadLines(os.Args[1])
	fmt.Println(solve(input))
}
