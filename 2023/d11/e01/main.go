package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets/intset"
	"fmt"
	"os"
	"strings"
)

func findExpandableCols(starMap []string) *intset.IntSet {
	ec := intset.New()
	for i := 0; i < len(starMap[0]); i++ {
		empty := true
		for j := 0; j < len(starMap); j++ {
			if starMap[j][i] != '.' {
				empty = false
			}
		}
		if empty {
			ec.Add(i)
		}
	}
	return ec
}

func findExpandableRows(starMap []string) *intset.IntSet {
	er := intset.New()
	for i := 0; i < len(starMap); i++ {
		empty := true
		for j := 0; j < len(starMap[0]); j++ {
			if starMap[i][j] != '.' {
				empty = false
			}
		}
		if empty {
			er.Add(i)
		}
	}
	return er
}

func findGalaxies(starMap []string) []coords2d.Coords2d {
	var galaxies []coords2d.Coords2d

	for i, line := range starMap {
		for j, cell := range line {
			if cell == '#' {
				galaxies = append(galaxies, coords2d.Coords2d{X: j, Y: i})
			}
		}
	}
	return galaxies
}

func sort(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func distance(left, right coords2d.Coords2d, ec, er *intset.IntSet) int {
	added := 0
	minX, maxX := sort(left.X, right.X)
	for i := minX; i < maxX; i++ {
		if ec.HasMember(i) {
			added++
		}
	}
	minY, maxY := sort(left.Y, right.Y)
	for i := minY; i < maxY; i++ {
		if er.HasMember(i) {
			added++
		}
	}
	return added + coords2d.Distance(left, right)
}

func solve(starMap []string) int {
	expandableCols := findExpandableCols(starMap)
	expandableRows := findExpandableRows(starMap)
	galaxies := findGalaxies(starMap)

	res := 0
	for i, left := range galaxies {
		for _, right := range galaxies[i+1:] {
			dist := distance(left, right, expandableCols, expandableRows)
			res += dist
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
