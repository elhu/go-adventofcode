package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

func solve(instructions []string) int {
	res := 0
	for _, inst := range instructions {
		hash := 0
		for _, c := range inst {
			hash += int(c)
			hash *= 17
			hash %= 256
		}
		res += hash
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	instructions := strings.Split(data, ",")
	fmt.Println(solve(instructions))
}
