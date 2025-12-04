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

func removeRolls(plan [][]byte) ([][]byte, int) {
	res := 0
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
					plan[i][j] = '.'
					res++
				}
			}
		}
	}
	return plan, res
}

func solve(plan [][]byte) int {
	plan = pad(plan, '.')
	res := 0
	for {
		var removed int
		plan, removed = removeRolls(plan)
		if removed == 0 {
			break
		}
		res += removed
	}
	return res
}

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	plan := bytes.Split(data, []byte("\n"))
	fmt.Println(solve(plan))
}
