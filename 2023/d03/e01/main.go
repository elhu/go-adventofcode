package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func pad(lines []string) []string {
	padded := make([]string, len(lines)+2)
	padded[0] = strings.Repeat(".", len(lines[0])+2)
	for i, line := range lines {
		padded[i+1] = "." + line + "."
	}
	padded[len(lines)+1] = strings.Repeat(".", len(lines[0])+2)
	return padded
}

func hasAdjacentSymbols(lines []string, i, start, end int) bool {
	for j := start - 1; j <= end; j++ {
		for x := i - 1; x <= i+1; x++ {
			if j >= start && j < end && x == i {
				continue
			}
			if (lines[x][j] <= '0' || lines[x][j] >= '9') && lines[x][j] != '.' {
				return true
			}
		}
	}
	return false
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	check(err)
	return n
}

func solve(lines []string) int {
	re := regexp.MustCompile(`\d+`)
	res := 0
	for i := 1; i < len(lines)-1; i++ {
		for _, match := range re.FindAllStringIndex(lines[i], -1) {
			if hasAdjacentSymbols(lines, i, match[0], match[1]) {
				res += atoi(lines[i][match[0]:match[1]])
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := pad(strings.Split(data, "\n"))
	fmt.Println(solve(lines))
}
