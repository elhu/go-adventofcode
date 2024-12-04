package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func solveLine(lines []string, startPos, vec coords2d.Coords2d) int {
	pos := startPos
	buffer := make([]coords2d.Coords2d, 0, 4)
	res := 0
	for pos.X >= 0 && pos.X < len(lines[0]) && pos.Y >= 0 && pos.Y < len(lines) {
		if bytes.ContainsAny([]byte{lines[pos.Y][pos.X]}, "XMAS") && lines[pos.Y][pos.X] != "XMAS"[len(buffer)] {
			buffer = buffer[:0]
		}
		if lines[pos.Y][pos.X] == "XMAS"[len(buffer)] {
			buffer = append(buffer, pos)
		}
		if len(buffer) == 4 {
			buffer = buffer[:0]
			res += 1
		}
		pos = coords2d.Add(pos, vec)
	}
	return res
}

func solve(lines []string) int {
	res := 0
	debug := make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		debug[i] = bytes.Repeat([]byte{'.'}, len(lines[0]))
	}
	for i := 0; i < len(lines); i++ {
		// Horizontal left-to-right
		res += solveLine(lines, coords2d.Coords2d{X: 0, Y: i}, coords2d.Coords2d{X: 1, Y: 0})
		// Horizontal right-to-left
		res += solveLine(lines, coords2d.Coords2d{X: len(lines[0]) - 1, Y: i}, coords2d.Coords2d{X: -1, Y: 0})
		// Diagonal top-left to bottom-right
		res += solveLine(lines, coords2d.Coords2d{X: 0, Y: i}, coords2d.Coords2d{X: 1, Y: 1})
		// Diagnonal bottom-left to top-right
		res += solveLine(lines, coords2d.Coords2d{X: 0, Y: i}, coords2d.Coords2d{X: 1, Y: -1})
		// Diagonal top-right to bottom-left
		res += solveLine(lines, coords2d.Coords2d{X: len(lines[0]) - 1, Y: i}, coords2d.Coords2d{X: -1, Y: 1})
		// Diagonal bottom-right to top-left
		res += solveLine(lines, coords2d.Coords2d{X: len(lines[0]) - 1, Y: i}, coords2d.Coords2d{X: -1, Y: -1})
	}
	for i := 0; i < len(lines[0]); i++ {
		// Vertical top-to-bottom
		res += solveLine(lines, coords2d.Coords2d{X: i, Y: 0}, coords2d.Coords2d{X: 0, Y: 1})
		// Vertical bottom-to-top
		res += solveLine(lines, coords2d.Coords2d{X: i, Y: len(lines) - 1}, coords2d.Coords2d{X: 0, Y: -1})
		// Don't double count corners
		if i > 0 && i < len(lines[0])-1 {
			// Diagonal top-left to bottom-right
			res += solveLine(lines, coords2d.Coords2d{X: i, Y: 0}, coords2d.Coords2d{X: 1, Y: 1})
			// Diagnonal top-right to bottom-left
			res += solveLine(lines, coords2d.Coords2d{X: i, Y: 0}, coords2d.Coords2d{X: -1, Y: 1})
			// Diagonal bottom-left to top-right
			res += solveLine(lines, coords2d.Coords2d{X: i, Y: len(lines) - 1}, coords2d.Coords2d{X: 1, Y: -1})
			// Diagonal bottom-right to top-left
			res += solveLine(lines, coords2d.Coords2d{X: i, Y: len(lines) - 1}, coords2d.Coords2d{X: -1, Y: -1})
		}
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	lines := strings.Split(data, "\n")
	fmt.Println(solve(lines))
}
