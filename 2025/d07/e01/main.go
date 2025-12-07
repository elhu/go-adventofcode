package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"adventofcode/utils/sets"
	"fmt"
	"os"
	"strings"
)

func solve(diag []string, start coords2d.Coords2d) int {
	beams := sets.New[coords2d.Coords2d]()
	beams.Add(start)
	split := 0
	for beams.Len() > 0 {
		for _, beam := range beams.Members() {
			beams.Remove(beam)
			if beam.Y == len(diag)-1 {
				break
			}
			if diag[beam.Y+1][beam.X] == '^' {
				split++
				beams.Add(coords2d.Coords2d{X: beam.X + 1, Y: beam.Y + 1})
				beams.Add(coords2d.Coords2d{X: beam.X - 1, Y: beam.Y + 1})
			} else {
				beams.Add(coords2d.Coords2d{X: beam.X, Y: beam.Y + 1})
			}
		}
	}
	return split
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	sLoc := coords2d.Coords2d{X: strings.Index(lines[0], "S"), Y: 0}
	fmt.Println(solve(lines, sLoc))
}
