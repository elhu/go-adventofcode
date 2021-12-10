package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

var openChars = "([{<"
var closeChars = ")]}>"

func findSyntaxError(line string) byte {
	var stack []byte
	var head byte
	for _, c := range []byte(line) {
		if idx := strings.IndexByte(closeChars, c); idx != -1 {
			oChar := openChars[idx]
			n := len(stack) - 1
			head, stack = stack[n], stack[:n]
			if oChar != head {
				return c
			}
		} else {
			stack = append(stack, c)
		}
	}
	return byte(0)
}

func solve(input []string) int {
	scores := map[byte]int{
		')': 3, ']': 57, '}': 1197, '>': 25137,
	}
	res := 0
	for _, l := range input {
		c := findSyntaxError(l)
		res += scores[c]
	}
	return res
}

func main() {
	input := files.ReadLines(os.Args[1])
	fmt.Println(solve(input))
}
