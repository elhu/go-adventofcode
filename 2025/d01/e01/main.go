package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(lines []string) int {
	dialPos := 50
	var res int
	for _, line := range lines {
		dir := line[0]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if dir == 'L' {
			dialPos -= clicks
		} else {
			dialPos += clicks
		}
		dialPos %= 100
		if dialPos == 0 {
			res += 1
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
