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

func mod(a, b int) int {
	return (a%b + b) % b
}

func lagrangeInterpolate(values [][2]int, toGuess int) int {
	var result float64
	for i := range values {
		term := float64(values[i][1])
		for j := range values {
			if i != j {
				term *= float64(toGuess-values[j][0]) / float64(values[i][0]-values[j][0])
			}
		}
		result += term
	}
	return int(result)
}

func bfs(lines [][]byte, startPos coords2d.Coords2d) int {
	toVisit := make([]*set.Set[coords2d.Coords2d], MAX_DEPTH+1)
	toVisit[0] = set.New[coords2d.Coords2d]()
	toVisit[0].Add(startPos)
	toVisit[1] = set.New[coords2d.Coords2d]()

	offset := MAX_DEPTH % len(lines)
	rounds := MAX_DEPTH/len(lines) + 1
	var sequence [][2]int

	for i := 0; i < MAX_DEPTH; i++ {
		toVisit[i+1] = set.New[coords2d.Coords2d]()
		for _, pos := range toVisit[i].Members() {
			for _, dir := range []coords2d.Coords2d{north, south, east, west} {
				next := coords2d.Add(pos, dir)
				modNext := coords2d.Coords2d{X: mod(next.X, len(lines[0])), Y: mod(next.Y, len(lines))}
				if lines[modNext.Y][modNext.X] == '.' || lines[modNext.Y][modNext.X] == 'S' {
					toVisit[i+1].Add(next)
				}
			}
		}
		if (i-offset)%len(lines) == 0 {
			sequence = append(sequence, [2]int{len(sequence) + 1, toVisit[i].Len()})
		}
		if len(sequence) == 3 {
			break
		}
	}
	return lagrangeInterpolate(sequence, rounds)
}

/*
* Since the input is square and repeats, the number of reachable plots is a polynomial function of the number of rounds,
* where each round is INPUT_LENGTH steps.
* MAX_DEPTH isn't a multiple of INPUT_LENGTH, so we need to start the sequence at offset MAX_DEPTH%INPUT_LENGTH.
 */

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

const MAX_DEPTH = 26501365

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	lines := bytes.Split(data, []byte("\n"))
	fmt.Println(solve(lines))
}
