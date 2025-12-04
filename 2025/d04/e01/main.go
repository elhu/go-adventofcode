package main

import (
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func pad(plan [][]byte, padChar byte) [][]byte {
	width := len(plan[0]) + 2
	paddedPlan := make([][]byte, 0, len(plan)+2)
	padLine := []byte(strings.Repeat(string(padChar), width))
	paddedPlan = append(paddedPlan, padLine)
	for _, line := range plan {
		line = append([]byte{padChar}, line...)
		line = append(line, padChar)
		paddedPlan = append(paddedPlan, line)
	}
	paddedPlan = append(paddedPlan, padLine)
	return paddedPlan
}

func solve(plan [][]byte) int {
	res := 0
	plan = pad(plan, '.')
	for i := 0; i < len(plan); i++ {
		for j := 0; j < len(plan[i]); j++ {
			if plan[i][j] == '@' {
				count := 0
				for k := i - 1; k <= i+1; k++ {
					for l := j - 1; l <= j+1; l++ {
						if plan[k][l] != '.' {
							count++
						}
					}
				}
				if count <= 4 {
					res++
				}
			}
		}
	}
	return res
}

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	plan := bytes.Split(data, []byte("\n"))
	fmt.Println(solve(plan))
}
