package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"fmt"
	"os"
	"strings"
)

var XMAS = []map[coords2d.Coords2d]byte{
	{
		coords2d.Coords2d{X: -1, Y: -1}: 'M',
		coords2d.Coords2d{X: 1, Y: -1}:  'M',
		coords2d.Coords2d{X: -1, Y: 1}:  'S',
		coords2d.Coords2d{X: 1, Y: 1}:   'S',
	}, {
		coords2d.Coords2d{X: -1, Y: -1}: 'M',
		coords2d.Coords2d{X: 1, Y: -1}:  'S',
		coords2d.Coords2d{X: -1, Y: 1}:  'M',
		coords2d.Coords2d{X: 1, Y: 1}:   'S',
	}, {
		coords2d.Coords2d{X: -1, Y: -1}: 'S',
		coords2d.Coords2d{X: 1, Y: -1}:  'M',
		coords2d.Coords2d{X: -1, Y: 1}:  'S',
		coords2d.Coords2d{X: 1, Y: 1}:   'M',
	}, {
		coords2d.Coords2d{X: -1, Y: -1}: 'S',
		coords2d.Coords2d{X: 1, Y: -1}:  'S',
		coords2d.Coords2d{X: -1, Y: 1}:  'M',
		coords2d.Coords2d{X: 1, Y: 1}:   'M',
	},
}

func checkXMAS(lines []string, pos coords2d.Coords2d) bool {
	for i := 0; i < len(XMAS); i++ {
		found := true
		for vec, v := range XMAS[i] {
			p := coords2d.Add(pos, vec)
			if lines[p.Y][p.X] != v {
				found = false
			}
		}
		if found {
			return true
		}
	}
	return false
}

func pad(lines []string) []string {
	res := make([]string, len(lines)+2)
	res[0] = strings.Repeat(".", len(lines[0])+2)
	res[len(res)-1] = strings.Repeat(".", len(lines[0])+2)
	for i := 0; i < len(lines); i++ {
		res[i+1] = "." + lines[i] + "."
	}
	return res
}

func solve(lines []string) int {
	res := 0
	for i := 1; i < len(lines)-1; i++ {
		for j := 0; j < len(lines[i])-1; j++ {
			if lines[i][j] == 'A' && checkXMAS(lines, coords2d.Coords2d{X: j, Y: i}) {
				res += 1
			}
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(pad(lines)))
}
