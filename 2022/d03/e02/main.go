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
	for i := 0; i < len(rucksacks); i += 3 {
		a, b, c := byteset.NewFromSlice([]byte(rucksacks[i])), byteset.NewFromSlice([]byte(rucksacks[i+1])), byteset.NewFromSlice([]byte(rucksacks[i+2]))
		sum += priority(a.Intersection(b).Intersection(c).Members()[0])
	}
	return sum
}

func main() {
	lines := files.ReadLines(os.Args[1])
	fmt.Println(solve(lines))
}
