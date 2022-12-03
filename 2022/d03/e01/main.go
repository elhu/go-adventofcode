package main

import (
	"adventofcode/utils/files"
	"adventofcode/utils/sets/byteset"
	"fmt"
	"os"
)

func priority(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	return int(c-'A') + 27
}

func solve(rucksacks []string) int {
	sum := 0
	for _, rs := range rucksacks {
		left, right := byteset.NewFromSlice([]byte(rs[0:len(rs)/2])), byteset.NewFromSlice([]byte(rs[len(rs)/2:]))
		sum += priority(left.Intersection(right).Members()[0])
	}
	return sum
}

func main() {
	lines := files.ReadLines(os.Args[1])
	fmt.Println(solve(lines))
}
