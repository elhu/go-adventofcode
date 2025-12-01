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
			for dialPos < 0 {
				res += 1
				dialPos += 100
			}
		} else {
			dialPos += clicks
			for dialPos >= 100 {
				res += 1
				dialPos -= 100
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
