package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
)

func directionalScore(x, y int, vectorX, vectorY int, heights [][]int, targetHeight int) int {
	score := 0
	x += vectorX
	y += vectorY
	for x >= 0 && x < len(heights[0]) && y >= 0 && y < len(heights) {
		score++
		if heights[y][x] >= targetHeight {
			break
		}
		x += vectorX
		y += vectorY
	}
	return score
}

func scenicScore(x, y int, heights [][]int) int {
	targetHeight := heights[y][x]
	score := 1
	score *= directionalScore(x, y, 0, -1, heights, targetHeight)
	score *= directionalScore(x, y, 0, 1, heights, targetHeight)
	score *= directionalScore(x, y, -1, 0, heights, targetHeight)
	score *= directionalScore(x, y, 1, 0, heights, targetHeight)

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
