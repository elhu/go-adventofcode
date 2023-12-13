package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func horizontalReflectionAt(row int, p []string) bool {
	for offset := 0; offset < min(row, len(p)-row); offset++ {
		if p[row-offset-1] != p[row+offset] {
			return false
		}
	}
	return true
}

func verticalReflectionAt(col int, p []string) bool {
	for offset := 0; offset < min(col, len(p[0])-col); offset++ {
		for i := 0; i < len(p); i++ {
			if p[i][col-offset-1] != p[i][col+offset] {
				return false
			}
		}
	}
	return true
}

func findReflection(p []string) int {
	for row := 1; row < len(p); row++ {
		if horizontalReflectionAt(row, p) {
			return row * 100
		}
	}
	for col := 1; col < len(p[0]); col++ {
		if verticalReflectionAt(col, p) {
			return col
		}
	}
	panic("WTF")
}

func solve(patterns [][]string) int {
	res := 0
	for _, p := range patterns {
		res += findReflection(p)
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	pdata := strings.Split(data, "\n\n")
	var patterns [][]string
	for _, p := range pdata {
		patterns = append(patterns, strings.Split(p, "\n"))
	}
	fmt.Println(solve(patterns))
}
