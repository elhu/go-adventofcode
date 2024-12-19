package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func compute(idx int, design string, patterns []string, cache map[string]int) int {
	res := 0
	if idx == len(design) {
		return 1
	}
	if val, ok := cache[design[idx:]]; ok {
		return val
	}
	for _, pattern := range patterns {
		if strings.HasPrefix(design[idx:], pattern) {
			res += compute(idx+len(pattern), design, patterns, cache)
			cache[design[idx:]] = res
		}
	}
	return res
}

func solve(designs []string, patterns []string) int {
	res := 0
	for _, design := range designs {
		c := compute(0, design, patterns, make(map[string]int))
		res += c
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	parts := strings.Split(data, "\n\n")

	patterns := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")
	fmt.Println(solve(designs, patterns))
}
