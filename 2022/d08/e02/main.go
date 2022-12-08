package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func scenicScore(x, y int, heights [][]int) int {
	targetHeight := heights[y][x]
	score := 1
	tmpScore := 0
	for i := y - 1; i >= 0; i-- {
		tmpScore++
		if heights[i][x] >= targetHeight {
			break
		}
	}
	score *= tmpScore

	tmpScore = 0
	for i := y + 1; i < len(heights); i++ {
		tmpScore++
		if heights[i][x] >= targetHeight {
			break
		}
	}
	score *= tmpScore

	tmpScore = 0
	for i := x - 1; i >= 0; i-- {
		tmpScore++
		if heights[y][i] >= targetHeight {
			break
		}
	}
	score *= tmpScore

	tmpScore = 0
	for i := x + 1; i < len(heights[0]); i++ {
		tmpScore++
		if heights[y][i] >= targetHeight {
			break
		}
	}
	score *= tmpScore

	return score
}

func solve(heights [][]int) int {
	res := 0
	for i := range heights {
		for j := range heights[i] {
			if score := scenicScore(j, i, heights); score > res {
				res = score
			}
		}
	}
	return res
}

func main() {
	data := files.ReadLines(os.Args[1])
	heights := make([][]int, len(data))
	for i, l := range data {
		heights[i] = make([]int, len(l))
		for j, c := range l {
			heights[i][j] = int(c - '0')
		}
	}
	fmt.Println(solve(heights))
}
