package main

import (
	"adventofcode/utils/coords/coords2d"
	"adventofcode/utils/files"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var directions = []coords2d.Coords2d{
	{X: 1, Y: 0},  // right
	{X: 0, Y: 1},  // down
	{X: -1, Y: 0}, // left
	{X: 0, Y: -1}, // up
}

func findStartTile(board [][]byte) coords2d.Coords2d {
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == '.' {
			return coords2d.Coords2d{X: i, Y: 0}
		}
	}
	panic("No starting tile found")
}

func oob(pos coords2d.Coords2d) bool {
	return pos.Y < 0 || pos.Y >= FACE_SIZE || pos.X < 0 || pos.X >= FACE_SIZE
}

type CoordTranslate struct {
	faceIndex, dirIndex int
	translate           func(coords2d.Coords2d) coords2d.Coords2d
}

var wrapsSample = [][]CoordTranslate{
	{
		{faceIndex: 5, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1, Y: FACE_SIZE - c.Y - 1}
		}},
		{faceIndex: 1, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: 0} }},
		{faceIndex: 2, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.Y, Y: 0} }},
		{faceIndex: 3, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1 - c.X, Y: 0} }},
	}, {
		{faceIndex: 5, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1 - c.Y, Y: 0} }},
		{faceIndex: 4, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: 0} }},
		{faceIndex: 2, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1, Y: c.Y} }},
		{faceIndex: 0, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: FACE_SIZE - 1} }},
	}, {
		{faceIndex: 1, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: c.Y} }},
		{faceIndex: 4, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: FACE_SIZE - 1 - c.X} }},
		{faceIndex: 3, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1, Y: c.Y} }},
		{faceIndex: 0, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: c.X} }},
	}, {
		{faceIndex: 2, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: c.Y} }},
		{faceIndex: 4, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1 - c.X, Y: FACE_SIZE - 1}
		}},
		{faceIndex: 5, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1 - c.Y, Y: FACE_SIZE - 1}
		}},
		{faceIndex: 0, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: 0} }},
	}, {
		{faceIndex: 5, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: c.Y} }},
		{faceIndex: 3, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1 - c.X, Y: FACE_SIZE - 1}
		}},
		{faceIndex: 2, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1 - c.Y, Y: FACE_SIZE - 1}
		}},
		{faceIndex: 1, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: FACE_SIZE - 1} }},
	}, {
		{faceIndex: 0, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1, Y: FACE_SIZE - 1 - c.Y}
		}},
		{faceIndex: 3, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: FACE_SIZE - 1 - c.X} }},
		{faceIndex: 4, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1, Y: c.Y} }},
		{faceIndex: 1, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1, Y: FACE_SIZE - 1 - c.X}
		}},
	},
}

var wraps = [][]CoordTranslate{
	{ // 0
		{faceIndex: 5, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: c.Y} }},
		{faceIndex: 1, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: 0} }},
		{faceIndex: 2, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: FACE_SIZE - 1 - c.Y} }},
		{faceIndex: 3, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: c.X} }},
	}, { // 1
		{faceIndex: 5, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.Y, Y: FACE_SIZE - 1} }},
		{faceIndex: 4, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: 0} }},
		{faceIndex: 2, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.Y, Y: 0} }},
		{faceIndex: 0, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: FACE_SIZE - 1} }},
	}, { // 2
		{faceIndex: 4, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: c.Y} }},
		{faceIndex: 3, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: 0} }},
		{faceIndex: 0, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: FACE_SIZE - 1 - c.Y} }},
		{faceIndex: 1, dirIndex: 0, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: 0, Y: c.X} }},
	}, { // 3
		{faceIndex: 4, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.Y, Y: FACE_SIZE - 1} }},
		{faceIndex: 5, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: 0} }},
		{faceIndex: 0, dirIndex: 1, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.Y, Y: 0} }},
		{faceIndex: 2, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: FACE_SIZE - 1} }},
	}, { // 4
		{faceIndex: 5, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1, Y: FACE_SIZE - 1 - c.Y}
		}},
		{faceIndex: 3, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1, Y: c.X} }},
		{faceIndex: 2, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1, Y: c.Y} }},
		{faceIndex: 1, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: FACE_SIZE - 1} }},
	}, { // 5
		{faceIndex: 4, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d {
			return coords2d.Coords2d{X: FACE_SIZE - 1, Y: FACE_SIZE - 1 - c.Y}
		}},
		{faceIndex: 1, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1, Y: c.X} }},
		{faceIndex: 0, dirIndex: 2, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: FACE_SIZE - 1, Y: c.Y} }},
		{faceIndex: 3, dirIndex: 3, translate: func(c coords2d.Coords2d) coords2d.Coords2d { return coords2d.Coords2d{X: c.X, Y: FACE_SIZE - 1} }},
	},
}

func wrap(pos coords2d.Coords2d, dirIndex, faceIndex int) (coords2d.Coords2d, int, int) {
	wrap := wraps[faceIndex][dirIndex]
	return wrap.translate(pos), wrap.dirIndex, wrap.faceIndex
}

func move(faces [6][][]byte, curr coords2d.Coords2d, dirIndex, faceIndex, distance int) (coords2d.Coords2d, int, int) {
	for i := 0; i < distance; i++ {
		next := coords2d.Add(curr, directions[dirIndex])
		nextDir := dirIndex
		nextFace := faceIndex
		if oob(next) {
			next, nextDir, nextFace = wrap(curr, dirIndex, faceIndex)
		}
		if faces[nextFace][next.Y][next.X] == '.' {
			curr = next
			dirIndex = nextDir
			faceIndex = nextFace
			visited = append(visited, V{c: curr, f: faceIndex, d: dirIndex})
		} else {
			return curr, dirIndex, faceIndex
		}
	}
	return curr, dirIndex, faceIndex
}

func mod(i, m int) int {
	res := i % m
	if res < 0 {
		return res + m
	}
	return res
}

type V struct {
	c coords2d.Coords2d
	f int
	d int
}

var dirs = [4]byte{'>', 'v', '<', '^'}

var visited []V

func solve(faces [6][][]byte, moves []byte) int {
	faceIndex := 0
	curr := findStartTile(faces[faceIndex])
	start := curr
	dirIndex := 0
	var distance int
	var direction byte
	visited = append(visited, V{c: curr, f: faceIndex, d: dirIndex})
	for len(moves) > 0 {
		fmt.Sscanf(string(moves), "%d", &distance)
		moves = moves[len(strconv.Itoa(distance)):]
		curr, dirIndex, faceIndex = move(faces, curr, dirIndex, faceIndex, distance)
		if len(moves) > 0 {
			direction = moves[0]
			if direction == 'L' {
				dirIndex = mod(dirIndex-1, len(directions))
			} else {
				dirIndex = mod(dirIndex+1, len(directions))
			}
			moves = moves[1:]
		}
	}
	board := stitchFaces(faces)
	for _, b := range board {
		fmt.Println(string(b))
	}
	fmt.Println("==")
	for _, v := range visited {
		board[FaceOffsets[v.f].Y*FACE_SIZE+v.c.Y][FaceOffsets[v.f].X*FACE_SIZE+v.c.X] = dirs[v.d]
	}
	board[FaceOffsets[0].Y*FACE_SIZE+start.Y][FaceOffsets[0].X*FACE_SIZE+start.X] = '&'
	for _, b := range board {
		fmt.Println(string(b))
	}
	return 1000*((FaceOffsets[faceIndex].Y*FACE_SIZE+curr.Y)+1) + 4*(FaceOffsets[faceIndex].X*FACE_SIZE+curr.X+1) + dirIndex
}

var FaceOffsetsSample = [6]coords2d.Coords2d{
	{X: 2, Y: 0},
	{X: 2, Y: 1},
	{X: 1, Y: 1},
	{X: 0, Y: 1},
	{X: 2, Y: 2},
	{X: 3, Y: 2},
}

var FaceOffsets = [6]coords2d.Coords2d{
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 0, Y: 2},
	{X: 0, Y: 3},
	{X: 1, Y: 2},
	{X: 2, Y: 0},
}

func parseFaces(board [][]byte) [6][][]byte {
	var faces [6][][]byte
	for i, fc := range FaceOffsets {
		faces[i] = make([][]byte, FACE_SIZE)
		for k := 0; k < FACE_SIZE; k++ {
			faces[i][k] = make([]byte, FACE_SIZE)
			copy(faces[i][k], board[fc.Y*FACE_SIZE+k][fc.X*FACE_SIZE:(fc.X+1)*FACE_SIZE])
		}
	}
	return faces
}

func stitchFaces(faces [6][][]byte) [][]byte {
	board := make([][]byte, FACE_SIZE*4)
	for i := range board {
		board[i] = make([]byte, FACE_SIZE*3)
		for j := range board[i] {
			board[i][j] = ' '
		}
	}
	for i, fc := range FaceOffsets {
		for j, values := range faces[i] {
			copy(board[fc.Y*FACE_SIZE+j][fc.X*FACE_SIZE:(fc.X+1)*FACE_SIZE], values)
		}
	}
	return board
}

const FACE_SIZE = 50

func main() {
	data := bytes.TrimRight(files.ReadFile(os.Args[1]), "\n")
	parts := bytes.Split(data, []byte("\n\n"))
	board := bytes.Split(parts[0], []byte("\n"))
	moves := parts[len(parts)-1]
	faces := parseFaces(board)
	fmt.Println(solve(faces, moves))
}
