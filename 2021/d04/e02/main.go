package main

import (
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	numbers [5][5]int
	seen    [5][5]bool
	hasWon  bool
}

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func parseGrids(data []string) []*Grid {
	res := make([]*Grid, 0)
	c := 0
	for c < len(data) {
		var gridNums [5][5]int
		for i := 0; i < 5; i++ {
			fields := strings.Fields(data[c+i])
			gridNums[i] = [5]int{atoi(fields[0]), atoi(fields[1]), atoi(fields[2]), atoi(fields[3]), atoi(fields[4])}
		}
		res = append(res, &Grid{numbers: gridNums, seen: [5][5]bool{}})
		c += 6
	}
	return res
}

func markNumber(d int, grid *Grid) (bool, int, int) {
	for i := range grid.numbers {
		for j := range grid.numbers[i] {
			if grid.numbers[i][j] == d {
				if !grid.hasWon {
					grid.seen[i][j] = true
				}
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

func hasWon(grid *Grid, y, x int) bool {
	seen := true
	for i := 0; i < len(grid.numbers); i++ {
		seen = seen && grid.seen[y][i]
	}
	if seen {
		return true
	}
	seen = true
	for i := 0; i < len(grid.numbers); i++ {
		seen = seen && grid.seen[i][x]
	}
	if seen {
		return true
	}
	return false
}

func computeScore(grid *Grid, d int) int {
	sum := 0
	for i := range grid.seen {
		for j := range grid.seen[i] {
			if !grid.seen[i][j] {
				sum += grid.numbers[i][j]
			}
		}
	}
	return sum * d
}

func solve(draws []int, grids []*Grid) int {
	var lastGid int
	var lastDraw int
	for _, d := range draws {
		for gid, g := range grids {
			if m, i, j := markNumber(d, g); m {
				if !g.hasWon && hasWon(g, i, j) {
					g.hasWon = true
					lastGid = gid
					lastDraw = d
				}
			}
		}
	}
	return computeScore(grids[lastGid], lastDraw)
}

func main() {
	data := files.ReadLines(os.Args[1])
	drawsStr := strings.Split(data[0], ",")
	draws := make([]int, len(drawsStr))
	for i, s := range drawsStr {
		draws[i] = atoi(s)
	}
	grids := parseGrids(data[2:])
	fmt.Println(solve(draws, grids))
}
