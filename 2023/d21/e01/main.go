package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	set "adventofcode/utils/sets"
	"bytes"
	"fmt"
	"os"
)

var (
	north = coords2d.Coords2d{X: 0, Y: -1}
	south = coords2d.Coords2d{X: 0, Y: 1}
	east  = coords2d.Coords2d{X: 1, Y: 0}
	west  = coords2d.Coords2d{X: -1, Y: 0}
)

func pad(lines [][]byte, val byte) [][]byte {
	padded := make([][]byte, len(lines)+2)
	padded[0] = bytes.Repeat([]byte{val}, len(lines[0])+2)
	for i, line := range lines {
		padded[i+1] = append([]byte{val}, line...)
		padded[i+1] = append(padded[i+1], val)
	}
	padded[len(lines)+1] = bytes.Repeat([]byte{val}, len(lines[0])+2)
	return padded
}

func bfs(lines [][]byte, startPos coords2d.Coords2d) int {
	toVisit := make([]*set.Set[coords2d.Coords2d], MAX_DEPTH+1)
	toVisit[0] = set.New[coords2d.Coords2d]()
	toVisit[0].Add(startPos)
	toVisit[1] = set.New[coords2d.Coords2d]()

	for i := 0; i < MAX_DEPTH; i++ {
		toVisit[i+1] = set.New[coords2d.Coords2d]()
		for _, pos := range toVisit[i].Members() {
			for _, dir := range []coords2d.Coords2d{north, south, east, west} {
				next := coords2d.Add(pos, dir)
				if lines[next.Y][next.X] == '.' || lines[next.Y][next.X] == 'S' {
					toVisit[i+1].Add(next)
				}
			}
		}
	}
	return toVisit[len(toVisit)-1].Len()
}

func solve(lines [][]byte) int {
	var startPos coords2d.Coords2d
	for y, line := range lines {
		if x := bytes.Index(line, []byte("S")); x != -1 {
			startPos = coords2d.Coords2d{X: x, Y: y}
			break
		}
	}
	return bfs(lines, startPos)
}

const MAX_DEPTH = 64

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	lines := bytes.Split(data, []byte("\n"))
	lines = pad(lines, '#')
	fmt.Println(solve(lines))
}
