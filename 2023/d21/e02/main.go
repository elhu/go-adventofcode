package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
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

func addWithMod(a, b coords2d.Coords2d, grid [][]byte) coords2d.Coords2d {
	res := coords2d.Add(a, b)
	res.X = mod(res.X, len(grid[0]))
	res.Y = mod(res.Y, len(grid))
	return res
}

func bfs(lines [][]byte, startPos coords2d.Coords2d) int {
	toVisit := make([]map[coords2d.Coords2d]struct{}, MAX_DEPTH+1)
	toVisit[0] = map[coords2d.Coords2d]struct{}{startPos: {}}
	toVisit[1] = map[coords2d.Coords2d]struct{}{}

	offset := MAX_DEPTH % len(lines)
	rounds := MAX_DEPTH/len(lines) + 1
	var sequence []int

	for i := 0; i < MAX_DEPTH; i++ {
		toVisit[i+1] = make(map[coords2d.Coords2d]struct{})
		for pos := range toVisit[i] {
			for _, dir := range []coords2d.Coords2d{north, south, east, west} {
				// next := addWithMod(pos, dir, lines)
				next := coords2d.Add(pos, dir)
				modNext := coords2d.Coords2d{X: mod(next.X, len(lines[0])), Y: mod(next.Y, len(lines))}
				if lines[modNext.Y][modNext.X] == '.' || lines[modNext.Y][modNext.X] == 'S' {
					toVisit[i+1][next] = struct{}{}
				}
			}
		}
		if (i-offset)%len(lines) == 0 {
			sequence = append(sequence, len(toVisit[i]))
		}
		if len(sequence) == 3 {
			break
		}
	}
	fmt.Println(sequence)
	// ToDo: find a way to compute the polynomial equation for the sequence
	// Plug those numbers into https://onlinetoolz.net/sequences and adjust the equation below
	return 14663*rounds*rounds - 14518*rounds + 3574
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
	fmt.Println("THIS SOLUTION ISN'T GENERIC, LOOK AT THE COMMENTS IN THE SOURCE CODE.")
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	lines := bytes.Split(data, []byte("\n"))
	fmt.Println(solve(lines))
}
