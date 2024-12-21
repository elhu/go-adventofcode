package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"os"
	"strconv"
	"strings"
)

var numKeyPad = map[coords2d.Coords2d]byte{
	{X: 0, Y: 3}: '7', {X: 1, Y: 3}: '8', {X: 2, Y: 3}: '9',
	{X: 0, Y: 2}: '4', {X: 1, Y: 2}: '5', {X: 2, Y: 2}: '6',
	{X: 0, Y: 1}: '1', {X: 1, Y: 1}: '2', {X: 2, Y: 1}: '3',
	{X: 0, Y: 0}: ' ', {X: 1, Y: 0}: '0', {X: 2, Y: 0}: 'A',
}

var dirKeyPad = map[coords2d.Coords2d]byte{
	{X: 0, Y: 1}: ' ', {X: 1, Y: 1}: '^', {X: 2, Y: 1}: 'A',
	{X: 0, Y: 0}: '<', {X: 1, Y: 0}: 'v', {X: 2, Y: 0}: '>',
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func findChar(pad map[coords2d.Coords2d]byte, char byte) coords2d.Coords2d {
	for k, v := range pad {
		if v == char {
			return k
		}
	}
	panic("WTF")
}

func cDiff(a, b coords2d.Coords2d) coords2d.Coords2d {
	return coords2d.Coords2d{X: a.X - b.X, Y: a.Y - b.Y}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getPadSeq(code string, startPos coords2d.Coords2d, pad map[coords2d.Coords2d]byte) string {
	var res string
	curr := startPos
	for _, c := range code {
		dest := findChar(pad, byte(c))
		delta := cDiff(dest, curr)

		var hMoves, vMoves string
		if delta.X != 0 {
			if delta.X > 0 {
				hMoves = strings.Repeat(">", abs(delta.X))
			} else {
				hMoves = strings.Repeat("<", abs(delta.X))
			}
		}
		if delta.Y != 0 {
			if delta.Y > 0 {
				vMoves = strings.Repeat("^", abs(delta.Y))
			} else {
				vMoves = strings.Repeat("v", abs(delta.Y))
			}
		}
		// Always move in one direction fully, then the other one
		trap := findChar(pad, ' ')
		if curr.Y == trap.Y && dest.X == trap.X { // Avoid hole coming from same row
			res += vMoves
			res += hMoves
		} else if curr.X == trap.X && dest.Y == trap.Y { // Avoid hole coming from same col
			res += hMoves
			res += vMoves
		} else if delta.X < 0 { // Move left before up (thanks reddit)
			res += hMoves
			res += vMoves
		} else { // Move down before left (thanks reddit)
			res += vMoves
			res += hMoves
		}
		res += "A"
		curr = dest
	}
	return res
}

func solveDirPads(targetSeq string, dirPadIdx int) int {
	seq := getPadSeq(targetSeq, findChar(dirKeyPad, 'A'), dirKeyPad)
	if dirPadIdx == 1 {
		return len(seq)
	}
	parts := strings.Split(seq, "A")
	count := 0
	for _, part := range parts[:len(parts)-1] {
		count += solveDirPads(part+"A", dirPadIdx-1)
	}
	return count
}

func solve(codes []string) int {
	res := 0
	for _, code := range codes {
		numSeq := getPadSeq(code, findChar(numKeyPad, 'A'), numKeyPad)
		count := solveDirPads(numSeq, 2)
		res += count * atoi(code[0:len(code)-1])
	}
	return res
}

func main() {
	data := strings.TrimRight(string(files.ReadFile(os.Args[1])), "\n")
	codes := strings.Split(data, "\n")
	println(solve(codes))
}
