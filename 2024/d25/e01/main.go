package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func parse(data []string) ([][]int, [][]int) {
	var keys, locks [][]int
	for _, block := range data {
		lines := strings.Split(block, "\n")

		block := make([]int, len(lines[0]))
		for j := 0; j < len(lines[0]); j++ {
			set := false
			for i := 1; i < len(lines); i++ {
				if lines[i][j] != lines[i-1][j] {
					set = true
					block[j] = i
					break
				}
			}
			if !set {
				block[j] = len(lines) - 1
			}
		}
		if lines[0] == strings.Repeat(".", len(lines[0])) {
			keys = append(keys, block)
		} else {
			locks = append(locks, block)
		}
	}
	return keys, locks
}

func solve(keys, locks [][]int) int {
	fmt.Println(keys)
	fmt.Println("----")
	fmt.Println(locks)
	res := 0
	for _, k := range keys {
		for _, l := range locks {
			fits := true
			for i := 0; i < len(k); i++ {
				if k[i] < l[i] {
					fits = false
					break
				}
			}
			if fits {
				res++
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n\n")
	keys, locks := parse(strings.Split(data, "\n\n"))
	fmt.Println(solve(keys, locks))
}
